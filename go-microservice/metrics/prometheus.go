package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	TotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "HTTP request latency",
		},
		[]string{"method", "endpoint"},
	)
)

func Init() {
	prometheus.MustRegister(TotalRequests, RequestDuration)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()
		next.ServeHTTP(w, r)
		RequestDuration.WithLabelValues(r.Method, r.URL.Path).
			Observe(time.Since(start).Seconds())
	})
}

func Handler() http.Handler {
	return promhttp.Handler()
}
