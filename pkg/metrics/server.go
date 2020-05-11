package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"github.com/wtty-fool/firestarter/pkg/server"
)

func NewMetricsServer(address string) *server.Server {
	logger := logrus.WithField("server", "metrics")

	handler := http.NewServeMux()
	handler.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			ErrorLog:          logger,
			EnableOpenMetrics: true,
			Timeout:           30 * time.Second,
		},
	))

	return server.NewServer(address, handler, logger)
}
