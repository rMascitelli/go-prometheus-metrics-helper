package main

import (
	"net/http"
	"time"
)

func main() {
	p := NewPrometheusClient()
	p.AddNewCounter("test_metric", "Random test metric")

	go func() {
		http.ListenAndServe(":2112", nil)
	}()
	for {
		p.IncrementMetric("test_metric")
		time.Sleep(time.Second * 2)
	}
}
