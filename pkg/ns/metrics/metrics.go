package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
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
