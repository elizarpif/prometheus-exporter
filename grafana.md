# Grafana

## Histogram
How to show histogram in Grafana?

This [article](https://opstrace.com/blog/grafana-histogram-howto) would be helpful
### heatmap

1. Promql
```
sum(rate(myapp_processed_time_bucket{}[1m])) by (le)
```

2. Set format to `Heatmap` and add `{{le}}` in Legend.

3. Set the data format to `Time series buckets`.

4. Set color map and hide zeros

5. Set axes unit to seconds

![1](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/grafana_histogram.png)

![2](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/grafana_histogram_settings.png)

What is this?

![3](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/grafana_histogram_legend.png)

_When you see a "dot" in the darkest red in the panel, then you can extract two additional data points:_

_the time of the observation — and you know that this observation actually stems from an observation period, which is 1 minute._
_a specific 'request processing duration' bucket (for ex, between 50 ms and 100 ms)
After reading out and rationalizing all three dimensions, we can try to put into words what this data point really means: you know that at said point in time, said bucket was hit with an average rate of 1.1 times per second, averaged across the observation period of 1 minute._


## Go metrics

```shell
## Number of goroutines that currently exist ({{pod}})
sum(go_goroutines{pod=~"${pod}"}) by (pod)

##-- go GC --
## Number of heap bytes when next garbage collection will take place ({{pod}})
sum(go_memstats_next_gc_bytes{pod=~"${pod}"}) by (pod)

## Number of bytes used for garbage collection system metadata  ({{pod}})
sum(go_memstats_gc_sys_bytes{pod=~"${pod}"}) by (pod)


## -- go memory --
## Number of heap bytes allocated and still in use | ({{pod}})
sum by(pod) (go_memstats_heap_alloc_bytes{pod=~"${pod}"})

## Number of heap bytes that are in use  ({{pod}})
sum by(pod) (go_memstats_heap_inuse_bytes{pod=~"${pod}"})

## Number of heap bytes waiting to be used  ({{pod}})
sum(go_memstats_heap_idle_bytes{pod=~"${pod}"}) by (pod)

## Number of bytes allocated and still in use  ({{pod}})
sum(go_memstats_alloc_bytes{pod=~"${pod}"}) by (pod)

## Number of stack bytes that are in use  ({{pod}})
sum(go_memstats_stack_inuse_bytes{pod=~"${pod}"}) by (pod)

```
