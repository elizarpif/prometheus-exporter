# prometheus-exporter
test prometheus and grafana

### go
Add simple metric
```go
opsProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"name", "id"})
```
And start server on 2112 port
```go
http.ListenAndServe(":2112", nil)
```

### Prometheus
1) Install prometheus using brew
```shell
brew install prometheus
```
2) Add our application in the prometheus config
```shell
vim /opt/homebrew/etc/prometheus.yml
```
```yaml
global:
  scrape_interval: 15s
  scrape_timeout: 10s
  evaluation_interval: 1m
scrape_configs:
- job_name: prometheus
  honor_timestamps: true
  scrape_interval: 15s
  scrape_timeout: 10s
  metrics_path: /metrics
  scheme: http
  follow_redirects: true
  static_configs:
  - targets:
    - localhost:9090
# our application
- job_name: myapp
  honor_timestamps: true
  scrape_interval: 10s
  scrape_timeout: 10s
  metrics_path: /metrics
  scheme: http
  follow_redirects: true
  static_configs:
  - targets:
    - localhost:2112
```
3) Start prometheus
```shell
brew services start prometheus
```
4) Open a link in the browser `localhost:9090`
5) Find our metric

![1](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/prometheus.png)


### Grafana
1) Install
```shell
brew install grafana
```
2) Edit the config, uncomment the next lines
```shell
vim /opt/homebrew/etc/grafana/grafana.ini
```
```yaml
    37 # The http port  to use
    38 http_port = 3000
    39
    40 # The public facing domain name used to access grafana from a browser
    41 domain = localhost
```
3) Start
```shell
/opt/homebrew/opt/grafana/bin/grafana-server --config /opt/homebrew/etc/grafana/grafana.ini --homepath /opt/homebrew/opt/grafana/share/grafana
```

4) Open a link in the browser `localhost:3000`

5) Add DataSource

![2](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/grafana_add_data_source.png)

After we see a successful test, add a dashboard and a dashboard

![3](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/grafana_panel.png)

## Push-gateway
1) Download push-gateway from docker and start in on port 9091
```shell
docker run -it -p 9091:9091 --rm prom/pushgateway
```
2) Add our application in the prometheus config
```shell
vim /opt/homebrew/etc/prometheus.yml
```
```yaml
global:
  scrape_interval: 15s
  scrape_timeout: 10s
  evaluation_interval: 1m
scrape_configs:
- job_name: prometheus
  honor_timestamps: true
  scrape_interval: 15s
  scrape_timeout: 10s
  metrics_path: /metrics
  scheme: http
  follow_redirects: true
  static_configs:
  - targets:
    - localhost:9090
# push gateway
- job_name: pushgateway
  honor_labels: true
  static_configs:
  - targets:
    - localhost:9091
```

Start gateway_pusher, go to `localhost:9091` and find our metric

![3](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/pushgateway.png)

And on prometheus as well

![3](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/prometheus2.png)
