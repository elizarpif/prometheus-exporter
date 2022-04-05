# Grafana

## Histogram
How to show histogram in Grafana?


### heatmap

```
sum(rate(myapp_processed_time_bucket{}[1m])) by (le)
```
![1](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/grafana_histogram.png)

![2](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/grafana_histogram_settings.png)
