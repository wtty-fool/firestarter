package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

const (
	httpAddress = ":9100"
)

func init() {
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}

func main() {
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))
	log.Infof("Starting to serve metrics at %s/metrics", httpAddress)
	log.Fatal(http.ListenAndServe(httpAddress, nil))
}
