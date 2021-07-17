package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type NewRequest struct {
	Lang  string `json:"lang,omitempty"`
	TTL   int64  `json:"ttl"`
	Level string `json:"level,omitempty"`
}

type NewResponse struct {
	ID     string `json:"id"`
	Image  string `json:"image"`
	Expire string `json:"expire"`
	Value  uint64 `json:"value"`
}

type SolveRequest struct {
	ID    string `json:"id"`
	Value uint64 `json:"value"`
}

func HTTPNew(c *fiber.Ctx, config *Config, storage *Storage) error {
	r := new(NewRequest)

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if r.TTL < 30 {
		r.TTL = 30
	} else if r.TTL >= 600 {
		r.TTL = 600
	}

	item := storage.NewItem(GetLevel(r.Level), r.Lang, r.TTL)
	image := GenerateCaptcha(item)

	response := NewResponse{ID: item.ID, Image: image, Expire: item.Expire.Format(time.RFC3339), Value: 0}

	if config.Test {
		response.Value = item.Value
	}

	PrometheusShowTotal.Inc()

	return c.JSON(response)
}

func HTTPSolve(c *fiber.Ctx, config *Config, storage *Storage) error {
	p := new(SolveRequest)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	valid := storage.Validate(p.ID, p.Value)

	if valid {
		return c.Status(200).JSON(true)
	}

	return c.Status(400).JSON(false)
}
