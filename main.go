package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

const (
	httpAddress = ":9100"
)

var (
	logger = log.New()
)

func init() {
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}

func main() {
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			ErrorLog:          logger,
			EnableOpenMetrics: true,
			Timeout:           30 * time.Second,
		},
	))
	logger.Infof("Starting to serve metrics at %s/metrics", httpAddress)
	logger.Fatal(http.ListenAndServe(httpAddress, nil))
}
