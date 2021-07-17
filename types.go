package main

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Storage struct {
	mu     sync.Mutex
	values map[string]uint64
	expire map[string]int64
}

type StorageItem struct {
	ID        string
	Value     uint64
	Level     int
	IntlValue string
	Language  string
	Expire    time.Time
}

type Config struct {
	Test bool
}

const (
	LevelMedium = 0
	LevelEasy   = 1
	LevelHard   = 2
)

var (
	// PrometheusStorageCount prometheus counter
	PrometheusStorageCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "restcaptcha_storage_count",
		Help: "The total number of storage items",
	})

	// PrometheusShowTotal prometheus counter
	PrometheusShowTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "restcaptcha_problem_request",
		Help: "The total number of request for show the captcha",
	})

	// PrometheusValidTotal prometheus counter
	PrometheusValidTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "restcaptcha_valid_request",
		Help: "The total number of request for success validate",
	})

	// PrometheusInValidTotal prometheus counter
	PrometheusInValidTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "restcaptcha_invalid_request",
		Help: "The total number of request for unsuccess validate",
	})
)
