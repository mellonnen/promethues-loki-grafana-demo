---
author: Arvid Gotthard (gotthard@kth.se)
date: January 2, 2006
paging: Slide %d / %d
---

# Monitoring a Golang web service using Prometheus, Loki, and Grafana.


---

# Stack
 
## Prometheus

- Time-series database for metrics.
- Built at SoundCloud in 2012.

## Loki

- Log aggregation system
- Inspired by Prometheus
- Developed by GrafanaLabs


## Grafana
- Data visualization/observability platform  
- Developed by GrafanaLabs


---

# System Architecture

```
┌─────┐ 
│ App │
└─────┘
```

- The app we want to monitor
  - For this demo it is a websocket server that powers a real-time chat app.


---

# System Architecture

```
┌─────┐   expose metrics   ┌──────────┐
│ App │ ─────────────────▶ │ /metrics │
└─────┘                    └──────────┘
```

- Inside the app we calculate metrics using the `prometheus/client_golang/` library
  - Metric types
    - Counter
    - Gauge
    - Histogram
    - Summary

- Expose the metrics on a http endpoint using `prometheus/client_golang/promhttp` library

---

# Cheat sheet


```js
  var (
    registerdClients = promauto.NewGauge(prometheus.GaugeOpts{
      Name: "registered_clients",
      Help: "Tracks the clients that are currently registered",
    })
  )
```


---

# System Architecture

```
┌─────┐   expose metrics   ┌──────────┐   scrape metrics  ┌────────────┐
│ App │ ─────────────────▶ │ /metrics │ ◀──────────────── │ Prometheus │  
└─────┘                    └──────────┘                   └────────────┘
```

- We configure Prometheus to scrape the metrics from the endpoint.

---

# System Architecture

```
┌─────┐   expose metrics   ┌──────────┐   scrape metrics  ┌────────────┐
│ App │ ─┬───────────────▶ │ /metrics │ ◀──────────────── │ Prometheus │  
└─────┘  │                 └──────────┘                   └────────────┘
         │                                                  
         │                 ┌──────┐ 
         └───────────────▶ │ Loki │
           push logs       └──────┘
```

- We push our logs from the app to Loki.

---

# System Architecture

```
┌─────┐   expose metrics   ┌──────────┐   scrape metrics  ┌────────────┐
│ App │ ─┬───────────────▶ │ /metrics │ ◀──────────────── │ Prometheus │  
└─────┘  │                 └──────────┘                   └────────────┘
         │                                                       ▲
         │                 ┌──────┐           ┌─────────┐        │
         └───────────────▶ │ Loki │ ◀───────  │ Grafana │ ───────┘  
           push logs       └──────┘   query   └─────────┘   query
```

- We configure Grafana to be able to query Loki and Prometheus as data sources.

---

# Configuring Prometheus, Loki and Grafana using `docker-compose`

```yaml
  grafana:
    container_name: grafana
    image: grafana/grafana:latest
    restart: unless-stopped
    user: "0"
    volumes:
      - "./grafana/data:/var/lib/grafana"
    ports:
      - 3000:3000
    environment:
      - GF_AUTH_DISABLE_LOGIN_FORM=true
      - GF_AUTH_ANONYMOUS_ENABLED=true
      - GF_AUTH_ANONYMOUS_ORG_ROLE=Admin
      - GF_SECURITY_ALLOW_EMBEDDING=true
    networks:
      - net

  loki:
    container_name: loki
    image: grafana/loki:2.3.0
    restart: unless-stopped
    ports:
      - 3100:3100
    volumes:
      - ./loki:/etc/loki
    command: -config.file=/etc/loki/loki-config.yml
    networks:
      - net

  prometheus:
    container_name: prometheus
    image: prom/prometheus
    restart: unless-stopped
    ports:
      - 9090:9090
    volumes:
      - ./prometheus:/etc/promethues/
    command: --config.file=/etc/promethues/prometheus-config.yml
    networks:
      - net
```

---

# Getting logs from container to Loki

---

# Getting logs from container to Loki

## Promtail

- Agent that lives on the host machine with the logs.
- It tails the log files and pushes them to Loki

---

# Getting logs from container to Loki

## Promtail

- Agent that lives on the host machine with the logs.
- It tails the log files and pushes them to Loki

## Docker plugin
- Provided by grafana

- Pushes logs from `stdout` and `stderr` of the container to loki

- Install by issuing:

```bash
$ docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions
```

- Configuration

```yaml
    logging:
      driver: loki
      options:
        loki-url: "http://localhost:3100/loki/api/v1/push"
```

---

# Thanks for listening!
