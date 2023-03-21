# prometheus-exporter
Тестовый прометей и графана

### go
Добавим такую метрику
```go
opsProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"name", "id"})
```
И запустим сервер на порту 2112
```go
http.ListenAndServe(":2112", nil)
```

### Prometheus
1) устанавливаем прометей через brew
```shell
brew install prometheus
```
2) добавляем в конфиг нашу приложеньку
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
# наша приложенька!!!
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
3) Запускаем Прометей
```shell
brew services start prometheus
```
4) В браузере открываем ссылку `localhost:9090`
5) Находим нашу метрику

![1](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/prometheus.png)


### Grafana
1) Устанавливаем
```shell
brew install grafana
```
2) Поправляем конфиг, раскомментируем следующие строчки
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
3) Запускаем
```shell
/opt/homebrew/opt/grafana/bin/grafana-server --config /opt/homebrew/etc/grafana/grafana.ini --homepath /opt/homebrew/opt/grafana/share/grafana
```

4) В браузере открываем ссылку `localhost:3000`

5) Добавляем DataSource

![2](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/grafana_add_data_source.png)

После того, как мы увидели, что тест успешен, добавляем дашборд и панель

![3](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/grafana_panel.png)

## Push-gateway
Для кронжоб, с которых не успевают собираться метрики.

1) Скачиваем push-gateway из докера и стартуем на порту 9091
```shell
docker run -it -p 9091:9091 --rm prom/pushgateway
```
2) Добавляем push-gateway в конфиг прометеуса, чтобы он мог читать метрики оттуда
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
Запускаем gateway_pusher, заходим на `localhost:9091` и находим нашу метрику

![3](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/pushgateway.png)

Находим в прометеусе эту же метрику

![3](https://github.com/elizarpif/prometheus-exporter/blob/develop/screens/prometheus2.png)
