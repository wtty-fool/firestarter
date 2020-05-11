package app

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/wtty-fool/firestarter/pkg/metrics"
	"github.com/wtty-fool/firestarter/pkg/server"
)

func home(w http.ResponseWriter, r *http.Request) {
	metrics.HomeVisitsTotal.Inc()
	fmt.Fprint(w, "Hello")
}

func NewAppServer(address string) *server.Server {
	logger := logrus.WithField("server", "app")

	handler := http.NewServeMux()
	handler.HandleFunc("/", home)

	return server.NewServer(address, handler, logger)
}
