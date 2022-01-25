package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

type sampleJson struct {
	Foo string `json:"foo"`
}

func TestHTTPEndpoint1(t *testing.T) {
	storage := NewStorage()
	config := Config{}

	config.ReturnValue = true
	config.TestImage = true

	app := fiber.New(fiber.Config{})

	app.Post("/new", func(c *fiber.Ctx) error {
		return HTTPNew(c, &config, storage)
	})

	app.Get("/test-image", func(c *fiber.Ctx) error {
		return HTTPNewTestImage(c, &config, storage)
	})

	app.Post("/solve", func(c *fiber.Ctx) error {
		return HTTPSolve(c, &config, storage)
	})

	req0 := httptest.NewRequest("GET", "/test-image", bytes.NewReader([]byte("")))
	resp0, _ := app.Test(req0)

	if resp0.StatusCode != 200 {
		t.Errorf("invalid response")
	}

	req1Body := NewRequest{
		Lang:  "fa",
		TTL:   10,
		Level: "EASY",
	}

	invalidReq1Body := sampleJson{
		Foo: "bar",
	}

	req1BodyJSON, _ := json.Marshal(req1Body)

	invalidReq1BodyJSON, _ := json.Marshal(invalidReq1Body)

	req1 := httptest.NewRequest("POST", "/new", bytes.NewReader(req1BodyJSON))
	req1.Header.Set("Content-Type", "application/json")

	req1Err := httptest.NewRequest("POST", "/new", bytes.NewReader(invalidReq1BodyJSON))

	resp1, _ := app.Test(req1)

	if resp1.StatusCode != 200 {
		t.Errorf("invalid response")
	}

	respErr1, _ := app.Test(req1Err)

	if respErr1.StatusCode != 422 {
		t.Errorf("invalid response")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp1.Body)
	str := buf.String()

	var newBody NewResponse

	json.Unmarshal([]byte(str), &newBody)

	req2Body := SolveRequest{
		ID:    newBody.ID,
		Value: newBody.Value,
	}

	req2BodyJSON, _ := json.Marshal(req2Body)

	req2 := httptest.NewRequest("POST", "/solve", bytes.NewReader(req2BodyJSON))
	req2.Header.Set("Content-Type", "application/json")

	req2Err := httptest.NewRequest("POST", "/solve", bytes.NewReader(req2BodyJSON))

	resp2, _ := app.Test(req2)

	resp2Err, _ := app.Test(req2Err)

	if resp2.StatusCode != 200 {
		t.Errorf("invalid response")
	}

	if resp2Err.StatusCode != 422 {
		t.Errorf("invalid response")
	}
}
func TestHTTPEndpoint2(t *testing.T) {
	storage := NewStorage()
	config := Config{}

	config.ReturnValue = false

	app := fiber.New(fiber.Config{})

	app.Post("/new", func(c *fiber.Ctx) error {
		return HTTPNew(c, &config, storage)
	})

	app.Post("/solve", func(c *fiber.Ctx) error {
		return HTTPSolve(c, &config, storage)
	})

	req1Body := NewRequest{
		Lang:  "fa",
		TTL:   86400,
		Level: "hard",
	}

	req1QualityBody := NewRequest{
		Lang:    "en",
		Quality: 100,
		TTL:     3600,
		Level:   "2",
	}

	invalidReq1Body := sampleJson{
		Foo: "bar",
	}

	req1BodyJSON, _ := json.Marshal(req1Body)

	invalidReq1BodyJSON, _ := json.Marshal(invalidReq1Body)

	req1QualityBodyJSON, _ := json.Marshal(req1QualityBody)
	req1Quality := httptest.NewRequest("POST", "/new", bytes.NewReader(req1QualityBodyJSON))
	req1Quality.Header.Set("Content-Type", "application/json")
	app.Test(req1Quality)

	req1 := httptest.NewRequest("POST", "/new", bytes.NewReader(req1BodyJSON))
	req1.Header.Set("Content-Type", "application/json")

	req1Err := httptest.NewRequest("POST", "/new", bytes.NewReader(invalidReq1BodyJSON))

	resp1, _ := app.Test(req1)

	if resp1.StatusCode != 200 {
		t.Errorf("invalid response")
	}

	respErr1, _ := app.Test(req1Err)

	if respErr1.StatusCode != 422 {
		t.Errorf("invalid response")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp1.Body)
	str := buf.String()

	var newBody NewResponse

	e := json.Unmarshal([]byte(str), &newBody)

	if e != nil || newBody.Value != 0 {
		t.Errorf("value must be zero")
	}

	req2Body := SolveRequest{
		ID:    newBody.ID,
		Value: 1,
	}

	req2BodyJSON, _ := json.Marshal(req2Body)

	req2 := httptest.NewRequest("POST", "/solve", bytes.NewReader(req2BodyJSON))
	req2.Header.Set("Content-Type", "application/json")

	resp2, _ := app.Test(req2)

	if resp2.StatusCode != 400 {
		t.Errorf("invalid response")
	}
}
