package monitoring

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Создание метрик для мониторинга запросов
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_requests_total",
			Help: "Total number of requests",
		},
		[]string{"method", "endpoint"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "app_request_duration_seconds",
			Help:    "Duration of requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	// Регистрация метрик
	prometheus.MustRegister(requestsTotal, requestDuration)
}

func StartPrometheusMetrics() {
	// Запуск HTTP-сервера для экспорта метрик
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Starting Prometheus metrics server on :2112/metrics")
	if err := http.ListenAndServe(":2112", nil); err != nil {
		log.Fatalf("Error starting Prometheus metrics server: %v", err)
	}
}
