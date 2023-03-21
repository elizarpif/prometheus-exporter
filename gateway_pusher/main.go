package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var names = []string{"malina-server", "strawberry", "cake"}
var results = []string{"error", "success"}
var times = []float64{1.23456, 3.45667, 2.3455, 1.234, 2.456, 4.6789, 10.23556}

func recordRandomHistogramMetric(vec *prometheus.HistogramVec) {
	name := names[rand.Intn(3)]
	t := times[rand.Intn(7)]
	res := results[rand.Intn(2)]

	vec.With(prometheus.Labels{"service_name": name, "operation_result": res}).Observe(t)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	ctx := context.Background()

	pp := NewPusher("metric", "localhost:9091")

	for i := 0; i < 20; i++ {
		recordRandomHistogramMetric(pp.histogram)

		err := pp.Push(ctx)
		if err != nil {
			fmt.Printf("Push: %v\n", err)
		}
	}

	time.Sleep(time.Hour)
}
