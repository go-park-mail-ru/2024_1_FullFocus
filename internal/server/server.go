package server

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.ServerConfig, r http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:        cfg.Port,
			Handler:     r,
			ReadTimeout: cfg.Timeout,
			IdleTimeout: cfg.IdleTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
