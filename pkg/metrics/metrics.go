package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	StartedTime = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "firestarter",
			Name:      "started_time",
			Help:      "Timestamp at which firestarter has started serving metrics",
		},
	)

	HomeVisitsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "firestarter",
			Name:      "home_visits",
			Help:      "Total visits to application homepage",
		},
	)

	HomeResponseTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "firestarter",
			Name:      "home_response_time_sec",
			Help:      "Time it took app to process a call to homepage (in seconds)",
			Buckets: []float64{
				1.0e-07, 2.5e-07, 5.0e-07, 7.5e-07, 9.0e-07,
				1.0e-06, 2.5e-06, 5.0e-06, 7.5e-06, 9.0e-06,
				1.0e-05, 2.5e-05, 5.0e-05, 7.5e-05, 9.0e-05,
			},
		},
	)
)

func init() {
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	prometheus.MustRegister(
		StartedTime,

		HomeVisitsTotal,
		HomeResponseTime,
	)

	StartedTime.SetToCurrentTime()
}
