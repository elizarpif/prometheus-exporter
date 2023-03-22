package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type pusher struct {
	name      string
	url       string
	done      chan bool
	histogram *prometheus.HistogramVec
	counter   *prometheus.CounterVec
	ticker    *time.Ticker
}

func NewPusher(recorderName, url string) *pusher {
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

	counterVec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "test",
		Name:      "test4",
		Help:      "",
	}, []string{"service_name", "operation_result"})
	prometheus.MustRegister(counterVec)

	return &pusher{
		name:      recorderName,
		url:       url,
		histogram: histogramV,
		counter:   counterVec,
		done:      make(chan bool),
		ticker:    time.NewTicker(time.Second * 5),
	}
}

func (p *pusher) Start() error {
	for {
		select {
		case <-p.done:
			return nil
		case <-p.ticker.C:
			err := p.add()
			if err != nil {
				return fmt.Errorf("add: %w", err)
			}
		}
	}
}

func (p *pusher) Close() error {
	p.ticker.Stop()
	close(p.done)

	err := p.add()
	if err != nil {
		return fmt.Errorf("close: %w", err)
	}

	return nil
}

func (p *pusher) add() error {
	err := push.New(p.url, p.name).
		Collector(p.histogram).
		Collector(p.counter).
		Add()
	if err != nil {
		return fmt.Errorf("add metrics: %w", err)
	}

	return nil
}

func (p *pusher) RecordLatency(result string, t float64) {
	p.histogram.With(prometheus.Labels{"service_name": p.name, "operation_result": result}).Observe(t)
}

func (p *pusher) RecordCounter(result string) {
	p.counter.With(prometheus.Labels{"service_name": p.name, "operation_result": result}).Inc()
}
