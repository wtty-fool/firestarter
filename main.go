package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
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
	log.Fatal(http.ListenAndServe(":9100", nil))
}
