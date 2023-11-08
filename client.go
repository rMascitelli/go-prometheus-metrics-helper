package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type PrometheusClient struct {
	gauges   map[string]*prometheus.GaugeVec
	counters map[string]*prometheus.CounterVec
}

func NewPrometheusClient() PrometheusClient {
	http.Handle("/metrics", promhttp.Handler())
	return PrometheusClient{
		gauges:   make(map[string]*prometheus.GaugeVec),
		counters: make(map[string]*prometheus.CounterVec),
	}
}

func (p *PrometheusClient) AddNewGauge(metricname string, description string) {
	if _, ok := p.gauges[metricname]; !ok {
		p.gauges[metricname] = promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: metricname, // metric name
				Help: description,
			},
			[]string{"hostname"},
		)
	} else {
		log.Printf("Metric %s already exists\n", metricname)
	}
}

func (p *PrometheusClient) AddNewCounter(metricname string, description string) {
	if _, ok := p.gauges[metricname]; !ok {
		p.counters[metricname] = promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: metricname, // metric name
				Help: description,
			},
			[]string{"hostname"},
		)
	} else {
		log.Printf("Metric %s already exists\n", metricname)
	}
}

func (p *PrometheusClient) IncrementCounter(metricname string, hostname string) {
	if counterVec, ok := p.counters[metricname]; ok {
		counterVec.WithLabelValues(hostname).Inc()
	} else {
		log.Printf("Metric %s doesnt exist\n", metricname)
	}
}

func (p *PrometheusClient) SetGaugeVal(metricname string, hostname string, gaugeval float64) {
	if gaugeVec, ok := p.gauges[metricname]; ok {
		gaugeVec.WithLabelValues(hostname).Set(gaugeval)
	} else {
		log.Printf("Metric %s doesnt exist\n", metricname)
	}
}
