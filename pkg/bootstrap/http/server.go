package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Option struct {
	Port    int
	Handler http.Handler
}

type BootstrapServer struct {
	*Option
	server *http.Server
}

func NewHttpServer(option *Option) *BootstrapServer {
	return &BootstrapServer{
		Option: option,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", option.Port),
			Handler: option.Handler,
		},
	}
}

func (s *BootstrapServer) Run() {
	go s.listenAndServe()
	s.listenCloseSignal()
	s.shutdown()
}

func (s *BootstrapServer) Close() {
	_ = s.server.Close()
}

func (s *BootstrapServer) listenAndServe() {
	log.Info().Msgf("BootstrapServer starting on port: %d", s.Option.Port)
	err := s.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panic().Err(err).Msg("Failed to start server")
	}
}

func (s *BootstrapServer) listenCloseSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-quit
}

func (s *BootstrapServer) shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Info().Msg("BootstrapServer shut downing...timeout 10 seconds")
	err := s.server.Shutdown(ctx)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panic().Err(err).Msg("Failed to shutdown server")
	}
	log.Info().Msg("BootstrapServer shutdown")
}
