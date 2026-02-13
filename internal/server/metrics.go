package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// HTTP request metrics
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: []float64{0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"method", "endpoint", "status_code"},
	)

	httpRequestSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_size_bytes",
			Help:    "HTTP request size in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 7),
		},
		[]string{"method", "endpoint"},
	)

	httpResponseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_size_bytes",
			Help:    "HTTP response size in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 7),
		},
		[]string{"method", "endpoint", "status_code"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(httpRequestSize)
	prometheus.MustRegister(httpResponseSize)
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    int64
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.written += int64(n)
	return n, err
}

// PrometheusMiddleware tracks HTTP requests and responses
func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap response writer to capture status code and size
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:    http.StatusOK,
		}

		// Record request size
		if r.ContentLength > 0 {
			httpRequestSize.WithLabelValues(r.Method, r.URL.Path).Observe(float64(r.ContentLength))
		}

		// Process request
		next.ServeHTTP(rw, r)

		// Calculate duration
		duration := time.Since(start).Seconds()

		// Record metrics
		statusCode := strconv.Itoa(rw.statusCode)
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, statusCode).Inc()
		httpRequestDuration.WithLabelValues(r.Method, r.URL.Path, statusCode).Observe(duration)
		httpResponseSize.WithLabelValues(r.Method, r.URL.Path, statusCode).Observe(float64(rw.written))
	})
}

// MetricsHandler returns the Prometheus metrics handler
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
