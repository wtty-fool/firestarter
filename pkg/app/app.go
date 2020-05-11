package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wtty-fool/firestarter/pkg/metrics"
	"github.com/wtty-fool/firestarter/pkg/server"
)

func responseLatency(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		metrics.HomeResponseTime.Observe(
			time.Now().Sub(start).Seconds(),
		)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	metrics.HomeVisitsTotal.Inc()
	fmt.Fprint(w, "Hello")
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func NewAppServer(address string) *server.Server {
	s := server.NewServer(address, "app")

	s.Router().Use(responseLatency)

	s.Router().HandleFunc("/", home)
	s.Router().HandleFunc("/favicon.ico", favicon)

	return s
}
