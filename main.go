package main

import (
	"github.com/sirupsen/logrus"

	"github.com/wtty-fool/firestarter/pkg/app"
	"github.com/wtty-fool/firestarter/pkg/metrics"
)

const (
	appAddress     = ":8000"
	metricsAddress = ":9100"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	metricsServer := metrics.NewMetricsServer(metricsAddress)
	go metricsServer.Listen()
	appServer := app.NewAppServer(appAddress)
	appServer.Listen()
}
