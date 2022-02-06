package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	prometheusStorageCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "restcaptcha_storage_count",
		Help: "The total number of storage items",
	})

	prometheusShowTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "restcaptcha_problem_request",
		Help: "The total number of request for show the captcha",
	})

	prometheusValidTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "restcaptcha_valid_request",
		Help: "The total number of request for success validate",
	})

	prometheusInValidTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "restcaptcha_invalid_request",
		Help: "The total number of request for unsuccess validate",
	})
)

func getPrometheusRegistry() *prometheus.Registry {
	promRegistry := prometheus.NewRegistry()
	promRegistry.MustRegister(prometheusStorageCount)
	promRegistry.MustRegister(prometheusShowTotal)
	promRegistry.MustRegister(prometheusValidTotal)
	promRegistry.MustRegister(prometheusInValidTotal)
	return promRegistry
}
