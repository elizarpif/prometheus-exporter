package main

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Pusher interface {
	Push(ctx context.Context) error
}

type pusher struct {
	serviceName string
	url         string
	histogram   *prometheus.HistogramVec
}

func NewPusher(label, url string) *pusher {
	histogramV := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "test",
			Name:      "message_handling_seconds",
			Help:      "Response time by module.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"service_name", "operation_result"},
	)
	prometheus.MustRegister(histogramV)

	return &pusher{
		serviceName: label,
		url:         url,
		histogram:   histogramV,
	}
}

var errPush = fmt.Errorf("push metrics")

func (p pusher) Push(ctx context.Context) error {
	err := push.New(p.url, p.serviceName).
		Collector(p.histogram).
		PushContext(ctx)
	if err != nil {
		return errPush
	}

	return nil
}
