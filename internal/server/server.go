package server

import (
	"critiqally/internal/config"
	"net/http"
)

type Server struct {
	cfg config.Config
	*http.Server
}

func New(cfg config.Config, handler http.Handler) Server {
	srv := &http.Server{
		Handler:      handler,
		Addr:         cfg.Address,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}
	return Server{
		cfg:    cfg,
		Server: srv,
	}
}
