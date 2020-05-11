package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/wtty-fool/firestarter/pkg/server"
)

func NewMetricsServer(address string) *server.Server {
	s := server.NewServer(address, "metrics")

	s.Router().Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			ErrorLog:          s.Logger(),
			EnableOpenMetrics: true,
			Timeout:           30 * time.Second,
		},
	))

	return s
}
