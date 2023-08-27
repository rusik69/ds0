package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "ds0_ns_requests_total",
		Help: "Total number of requests to ds0 ns",
	})
)

func init() {
	prometheus.MustRegister(Counter)
}

// PrometheusHandler returns the prometheus handler.
func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
