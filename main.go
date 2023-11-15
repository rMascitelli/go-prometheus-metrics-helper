package main

import (
	client "github.com/rMascitelli/go-prometheus-metrics-helper/client"
	"log"
	"net/http"
	"time"
)

func main() {
	p := client.NewPrometheusClient()
	servicename := "randomServiceName"
	p.AddNewCounter("test_counter", "Testing out a counter")
	p.AddNewGauge("test_gauge", "Testing out a gauge")

	go func() {
		log.Println("Hosting metrics at localhost:2112")
		http.ListenAndServe(":2112", nil)
	}()
	for {
		randfloat := float64(time.Now().Unix()) / 13
		p.IncrementCounter("test_counter", servicename)
		p.SetGaugeVal("test_gauge", servicename, randfloat)
		time.Sleep(time.Second * 2)
	}
}
