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

	HomeResponseTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "firestarter",
			Name:      "home_response_time",
			Help:      "Time it took app to process a call to homepage",
			Buckets:   []float64{0.001, 0.01, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0},
		},
	)

	allMetrics = []prometheus.Collector{
		HomeVisitsTotal,
		HomeResponseTime,
	}
)

func init() {
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	prometheus.MustRegister(allMetrics...)
}
