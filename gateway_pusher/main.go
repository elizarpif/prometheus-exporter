package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var names = []string{"malina-server", "strawberry", "cake"}
var results = []string{"error", "success"}
var times = []float64{1.23456, 3.45667, 2.3455, 1.234, 2.456, 4.6789, 10.23556}

func recordMetric(p *pusher) {
	t := times[rand.Intn(7)]
	res := results[rand.Intn(2)]

	p.RecordLatency(res, t)
	p.RecordCounter(res)

	fmt.Printf("recorded:\n histogram result=%v, time=%v\n counter result=%v", res, t, res)
}

func main() {
	pp := NewPusher("cake", "localhost:9091")

	wt := sync.WaitGroup{}
	wt.Add(1)

	go func() {
		err := pp.Start()
		if err != nil {
			panic(err)
		}
		wt.Done()
	}()

	for i := 0; i < 20; i++ {
		recordMetric(pp)
		time.Sleep(time.Second)
	}

	err := pp.Close()
	if err != nil {
		panic(err)
	}

	wt.Wait()
}
