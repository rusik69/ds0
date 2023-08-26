package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "ds0_node_requests_total",
		Help: "Total number of requests to ds0 node",
	})
)

func init() {
	prometheus.MustRegister(Counter)
}
