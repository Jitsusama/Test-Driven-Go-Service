// Package server is responsible for exposing an HTTP server interface that has handlers associated
// with accepting weather requests and returning translated responses.
package server

import (
	"context"
	"fmt"
	"net/http"
	"playing-around/pkg/translator"
	"strings"
)

// Server represents an HTTP server.
type Server struct {
	port   int
	ctx    context.Context
	server http.Server
	trans  translator.Translator
}

// Create a new server.
func Create(port int, trans translator.Translator) *Server {
	return &Server{port: port, ctx: context.Background(), trans: trans}
}

// Start the server.
func (s *Server) Start() error {
	handler := s.createHandler()
	s.server = s.createServer(handler)
	return s.server.ListenAndServe()
}

// Stop the server.
func (s *Server) Stop() error {
	return s.server.Shutdown(s.ctx)
}

func (s *Server) createHandler() *http.ServeMux {
	handler := http.NewServeMux()
	handler.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")

		city := strings.TrimPrefix(r.URL.Path, "/weather/")
		temp := s.trans.RetrieveWeather(city)

		_, _ = w.Write([]byte(temp))
	})
	return handler
}

func (s *Server) createServer(handler *http.ServeMux) http.Server {
	return http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: handler,
	}
}
