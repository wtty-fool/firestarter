package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	address string
	handler http.Handler
	logger  *logrus.Entry
}

func (s *Server) Listen() {
	s.logger.Infof("Starting to listen and serve at %s", s.address)
	s.logger.Fatal(http.ListenAndServe(s.address, s.handler))
}

func NewServer(address string, handler http.Handler, logger *logrus.Entry) *Server {
	return &Server{
		address: address,
		handler: handler,
		logger:  logger,
	}
}
