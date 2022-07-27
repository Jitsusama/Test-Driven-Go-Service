package server

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	port   int
	ctx    context.Context
	server http.Server
}

func Create(port int) *Server {
	return &Server{port: port, ctx: context.Background()}
}

func (s *Server) Start() error {
	s.server = http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: nil,
	}
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Shutdown(s.ctx)
}
