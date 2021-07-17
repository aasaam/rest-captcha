package main

import (
	"flag"
	"strings"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func interval(storage *Storage) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				storage.CleanUp()
				PrometheusStorageCount.Set(float64(storage.Count()))
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func main() {
	storage := NewStorage()
	config := Config{}

	var baseURL string
	flag.StringVar(&baseURL, "base-url", "/", "Base URL for routes")
	baseURL = strings.TrimRight(baseURL, "/")

	var listen string
	flag.StringVar(&listen, "listen", "0.0.0.0:4000", "Application listen address")

	config.Test = *flag.Bool("test", false, "Test m")

	promRegistry := prometheus.NewRegistry()
	promRegistry.MustRegister(PrometheusStorageCount)
	promRegistry.MustRegister(PrometheusInValidTotal)
	promRegistry.MustRegister(PrometheusValidTotal)

	handler := promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{})

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               true,
		UnescapePath:          true,
		CaseSensitive:         true,
		StrictRouting:         true,
		BodyLimit:             1 * 512,

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			ctx.Status(code).SendString(err.Error())

			return nil
		},
	})

	app.Get("/metrics", adaptor.HTTPHandler(handler))

	app.Post(baseURL+"/new", func(c *fiber.Ctx) error {
		return HTTPNew(c, &config, storage)
	})

	app.Post(baseURL+"/solve", func(c *fiber.Ctx) error {
		return HTTPSolve(c, &config, storage)
	})

	interval(storage)

	e := app.Listen(listen)
	if e != nil {
		panic(e.Error())
	}
}
