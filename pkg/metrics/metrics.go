package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var ()

func init() {
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}

type MetricsServer struct {
	address string
	handler *http.ServeMux
	logger  *logrus.Entry
}

func (m *MetricsServer) Listen() {
	m.logger.Infof("Starting to serve metrics at %s/metrics", m.address)
	m.logger.Fatal(http.ListenAndServe(m.address, m.handler))
}

func NewMetricsServer(address string) *MetricsServer {
	logger := logrus.WithField("logger", "metrics")

	handler := http.NewServeMux()
	handler.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			ErrorLog:          logger,
			EnableOpenMetrics: true,
			Timeout:           30 * time.Second,
		},
	))

	return &MetricsServer{
		address: address,
		handler: handler,
		logger:  logger,
	}
}
