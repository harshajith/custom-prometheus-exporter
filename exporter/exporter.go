package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define a custom counter metric
	myCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "my_custom_counter",
			Help: "This is my custom counter",
		},
	)

	// Define a custom gauge metric
	myGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "my_custom_gauge",
			Help: "This is my custom gauge",
		},
	)
)

func init() {
	// Register custom metrics with the Prometheus client library
	prometheus.MustRegister(myCounter)
	prometheus.MustRegister(myGauge)
}

func recordMetrics() {
	go func() {
		for {
			myCounter.Inc()                   // Increment the counter
			myGauge.Set(rand.Float64() * 100) // Set the gauge to a random value between 0 and 100
			time.Sleep(2 * time.Second)       // Sleep for 2 seconds
		}
	}()
}

func main() {
	// Start recording metrics
	recordMetrics()

	// Expose metrics at `/metrics` endpoint
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
