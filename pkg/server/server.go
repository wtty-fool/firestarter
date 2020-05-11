package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	address string
	router  *mux.Router
	logger  *logrus.Entry
}

// loggingMiddleware is a used to log all incoming requests.
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.WithFields(
			logrus.Fields{
				"remote_addr": r.RemoteAddr,
				"method":      r.Method,
				"uri":         r.RequestURI,
			},
		).Info("-")
		next.ServeHTTP(w, r)
	})
}

func (s *Server) Listen() {
	s.logger.Infof("Starting to listen and serve at %s", s.address)
	s.logger.Fatal(http.ListenAndServe(s.address, s.router))
}

func (s *Server) Logger() *logrus.Entry {
	return s.logger
}

func (s *Server) Router() *mux.Router {
	return s.router
}

func NewServer(address, name string) *Server {
	s := &Server{
		address: address,
		router:  nil,
		logger:  logrus.WithField("server", name),
	}

	r := mux.NewRouter()
	r.Use(s.loggingMiddleware)
	s.router = r

	return s
}
