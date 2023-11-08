package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	p := NewPrometheusClient()
	hostname := "randomServiceName"
	p.AddNewCounter("test_counter", "Testing out a counter")
	p.AddNewGauge("test_gauge", "Testing out a gauge")

	go func() {
		log.Println("Hosting metrics at localhost:2112")
		http.ListenAndServe(":2112", nil)
	}()
	for {
		randfloat := float64(time.Now().Unix()) / 13
		p.IncrementCounter("test_counter", hostname)
		p.SetGaugeVal("test_gauge", hostname, randfloat)
		time.Sleep(time.Second * 2)
	}
}
