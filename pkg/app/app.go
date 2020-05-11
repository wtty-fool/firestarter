package app

import (
	"fmt"
	"net/http"

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
	s := server.NewServer(address, "app")

	s.Router().HandleFunc("/", home)
	s.Router().HandleFunc("/favicon.ico", favicon)

	return s
}
