package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func interval(storage *storage) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				storage.cleanUp()
				prometheusStorageCount.Set(float64(storage.count()))
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func main() {
	var baseURL string
	flag.StringVar(&baseURL, "base-url", "/", "Base URL for routes")

	var username string
	flag.StringVar(&username, "auth-username", "", "Basic authentication username")

	var password string
	flag.StringVar(&password, "auth-password", "", "Basic authentication password")

	var listen string
	flag.StringVar(&listen, "listen", "0.0.0.0:4000", "Application listen address")

	returnValue := flag.Bool("return-value", false, "Return value on generation")

	testImage := flag.Bool("test-image", false, "Expose /test-image for testing image")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	storage := newStorage()
	config := config{}

	config.returnValue = *returnValue
	config.testImage = *testImage

	baseURL = strings.TrimRight(baseURL, "/")

	promRegistry := getPrometheusRegistry()

	handler := promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{})

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               false,
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

	if username != "" && password != "" {
		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				username: password,
			},
			Realm: "REST Captcha",
		}))
	}

	app.Get("/metrics", adaptor.HTTPHandler(handler))

	app.Post(baseURL+"/new", func(c *fiber.Ctx) error {
		return httpNew(c, &config, storage)
	})

	app.Post(baseURL+"/solve", func(c *fiber.Ctx) error {
		return httpSolve(c, &config, storage)
	})

	if config.testImage {
		app.Get(baseURL+"/test-image", func(c *fiber.Ctx) error {
			return httpNewTestImage(c, &config, storage)
		})
	}

	interval(storage)

	e := app.Listen(listen)
	if e != nil {
		panic(e.Error())
	}
}
