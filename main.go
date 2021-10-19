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
	sorterNames := []string{"Alice", "Peter", "Vlad"}
	warehouseIDs := []int64{123456, 345667, 23455}

	go func() {
		for {
			sorterName := sorterNames[rand.Intn(3)]
			wareHouseID := warehouseIDs[rand.Intn(3)]

			opsProcessed.With(prometheus.Labels{"name": sorterName, "id": strconv.Itoa(int(wareHouseID))}).Inc()
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"name", "id"})
)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	logrus.Info("started localhost:2112")
	http.ListenAndServe(":2112", nil)
}
