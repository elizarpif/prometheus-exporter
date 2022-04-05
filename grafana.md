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