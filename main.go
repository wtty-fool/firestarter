package main

import (
	"github.com/sirupsen/logrus"

	"github.com/wtty-fool/firestarter/pkg/metrics"
)

const (
	metricsAddress = ":9100"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	metricsServer := metrics.NewMetricsServer(metricsAddress)
	go metricsServer.Listen()
}
