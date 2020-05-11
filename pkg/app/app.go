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

func favicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func NewAppServer(address string) *server.Server {
	logger := logrus.WithField("server", "app")

	handler := http.NewServeMux()
	handler.HandleFunc("/", home)
	handler.HandleFunc("/favicon.ico", favicon)

	return server.NewServer(address, handler, logger)
}
