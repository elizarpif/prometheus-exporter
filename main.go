package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func recordMetrics() {
	names := []string{"Alice", "Peter", "Vlad"}
	someIDs := []int64{123456, 345667, 23455}

	go func() {
		for {
			name := names[rand.Intn(3)]
			someID := someIDs[rand.Intn(3)]

			opsProcessed.With(prometheus.Labels{"name": name, "id": strconv.Itoa(int(someID))}).Inc()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			timeStart := time.Now()

			tm := rand.Intn(1000)
			time.Sleep(time.Millisecond * time.Duration(tm))

			histogram.Observe(float64(time.Since(timeStart)))
		}
	}()
}

var (
	// counter increments only
	opsProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"name", "id"})

	// histogram can be helpful for timing
	histogram = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:        "myapp_processed_time",
		Help:        "The average time (seconds) of the process",
		Buckets:     []float64{0.001, 0.002, 0.005, 0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2},
	})
)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	logrus.Info("started localhost:2112")
	http.ListenAndServe(":2112", nil)
}
