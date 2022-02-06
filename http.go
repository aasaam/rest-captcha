package main

import (
	"encoding/base64"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type newRequest struct {
	Lang    string `json:"lang,omitempty"`
	TTL     int64  `json:"ttl"`
	Quality int    `json:"quality"`
	Level   string `json:"level,omitempty"`
}

type newResponse struct {
	ID     string `json:"id"`
	Image  string `json:"image"`
	Expire string `json:"expire"`
	Value  uint64 `json:"value"`
}

type solveRequest struct {
	ID    string `json:"id"`
	Value uint64 `json:"value"`
}

func httpNewTestImage(c *fiber.Ctx, config *config, storage *storage) error {
	r := new(newRequest)

	r.Lang = c.Query("lang", "en")
	quality, _ := strconv.Atoi(c.Query("q", "0"))

	r.Level = c.Query("level", "0")

	item := storage.newItem(getLevel(r.Level), r.Lang, r.TTL)
	image := generateCaptcha(item, quality)

	imageByte, _ := base64.StdEncoding.DecodeString(image)
	c.Set("X-Value", strconv.Itoa(int(item.value)))
	c.Set("Content-Type", "image/jpeg")

	return c.Send(imageByte)
}

func httpNew(c *fiber.Ctx, config *config, storage *storage) error {
	r := new(newRequest)

	if err := c.BodyParser(r); err != nil {
		return err
	}

	quality := 30
	if r.Quality > 0 {
		quality = minMaxDefault(r.Quality, 1, 95)
	}

	r.TTL = minMaxDefault64(r.TTL, 30, 600)

	item := storage.newItem(getLevel(r.Level), r.Lang, r.TTL)
	image := generateCaptcha(item, quality)

	response := newResponse{ID: item.id, Image: "data:image/jpeg;base64," + image, Expire: item.expire.Format(time.RFC3339), Value: 0}

	if config.returnValue {
		response.Value = item.value
	}

	prometheusShowTotal.Inc()

	return c.JSON(response)
}

func httpSolve(c *fiber.Ctx, config *config, storage *storage) error {
	r := new(solveRequest)

	if err := c.BodyParser(r); err != nil {
		return err
	}

	valid := storage.validate(r.ID, r.Value)

	if valid {
		return c.Status(200).JSON(true)
	}

	return c.Status(400).JSON(false)
}
