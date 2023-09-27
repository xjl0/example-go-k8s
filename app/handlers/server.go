package handlers

import (
	"context"
	"net/http"
	"time"
)

// Server struct
type Server struct {
	httpServer *http.Server
}

// Run http server
func (s *Server) Run(host, port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         host + ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown http server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
