// push_to_gateway.go pushes a timestamp to prom/pushgateway operating at
// :9091. To see its result, you have to have a pushgateway and prometheus
// running. This setup simulates metric reporting for jobs/cronjobs.
package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/sirupsen/logrus"
)

func main() {
	g := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "firestarter",
			Name:      "scripts_push_to_gateway_time",
			Help:      "Timestamp at which push_to_gateway script ran",
		},
	)

	registry := prometheus.NewRegistry()
	registry.MustRegister(g)

	pusher := push.New(
		"http://localhost:9091",
		"firestarter_push_to_gateway",
	).Gatherer(registry)
	g.SetToCurrentTime()

	if err := pusher.Push(); err != nil {
		logrus.Fatal(err)
	}
}
