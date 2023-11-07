package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type PrometheusClient struct {
	Metrics map[string]prometheus.Counter
}

func NewPrometheusClient() PrometheusClient {
	http.Handle("/metrics", promhttp.Handler())
	return PrometheusClient{
		Metrics: make(map[string]prometheus.Counter),
	}
}

func (p *PrometheusClient) AddNewCounter(name string, desc string) {
	newCounterPtr := promauto.NewCounter(
		prometheus.CounterOpts{
			Name: name, // metric name
			Help: desc,
		},
	)
	fmt.Printf("metrics: %+v\n", p.Metrics)
	p.Metrics[name] = newCounterPtr
}

func (p *PrometheusClient) IncrementMetric(name string) {
	p.Metrics[name].Inc()
}
