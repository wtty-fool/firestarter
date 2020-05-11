package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HomeVisitsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "firestarter",
			Name:      "home_visits_total",
			Help:      "Total visits to application homepage",
		},
	)

	allMetrics = []prometheus.Collector{
		HomeVisitsTotal,
	}
)

func init() {
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	prometheus.MustRegister(allMetrics...)
}
