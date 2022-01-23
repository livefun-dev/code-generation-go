package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//go:generate go run ../cmd/generator/main.go

func recordMetrics() {
	go func() {
		for {
			//+prom:metric:counter name:myapp_processed_ops_total
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func recordMetrics2() {
	go func() {
		for {
			//+prom:metric:gauge name:myapp_processed_ops_total_second
			mySecondMetric.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()

	fmt.Printf("🚀 Running metric server: http://localhost:2112/metrics \n")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
