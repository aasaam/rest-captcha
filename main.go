package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
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

func runCaptchaServer(c *cli.Context) error {

	config := config{}
	config.returnValue = c.Bool("return-value")
	config.testImage = c.Bool("test-image")
	username := c.String("auth-username")
	password := c.String("auth-password")

	storage := newStorage(!config.returnValue)

	baseURL := strings.TrimRight(c.String("base-url"), "/")

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

	if !config.returnValue {
		interval(storage)
	}

	return app.Listen(c.String("listen"))
}

func main() {
	app := cli.NewApp()
	app.Usage = "REST Captcha"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:   "run",
			Usage:  "Run captcha server",
			Action: runCaptchaServer,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "base-url",
					Usage:    "Base URL for routes",
					Value:    "/",
					Required: false,
					EnvVars:  []string{"ASM_BASE_URL"},
				},
				&cli.StringFlag{
					Name:     "auth-username",
					Usage:    "Basic authentication username",
					Value:    "",
					Required: false,
					EnvVars:  []string{"ASM_AUTH_USERNAME"},
				},
				&cli.StringFlag{
					Name:     "auth-password",
					Usage:    "Basic authentication password",
					Value:    "",
					Required: false,
					EnvVars:  []string{"ASM_AUTH_PASSWORD"},
				},
				&cli.StringFlag{
					Name:     "listen",
					Usage:    "Application listen http ip:port address",
					Value:    "127.0.0.1:4000",
					Required: false,
					EnvVars:  []string{"ASM_AUTH_PASSWORD"},
				},
				&cli.BoolFlag{
					Name:     "return-value",
					Usage:    "Application listen http ip:port address",
					Value:    false,
					Required: false,
					EnvVars:  []string{"ASM_RETURN_VALUE"},
				},
				&cli.BoolFlag{
					Name:     "test-image",
					Usage:    "Expose /test-image for testing image",
					Value:    false,
					Required: false,
					EnvVars:  []string{"ASM_TEST_IMAGE"},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
