package main

import (
	"context"
	"critiqally/internal/config"
	"critiqally/internal/router"
	"critiqally/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("initializing configuration")
	cfg := config.New()

	if cfg.PrettyLogging {
		log.Logger = log.Output(zerolog.NewConsoleWriter())
	}

	log.Info().Msg("initializing router")
	r := router.New(cfg)

	log.Info().Msg("initializing server")
	s := server.New(cfg, r)

	// Channel to listen for termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info().Msg("starting server")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("server failed to start")
		}
	}()

	<-quit
	log.Info().Msg("shutting down server...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server forced to shutdown")
	}

	log.Info().Msg("server exiting")
}
