package handler

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func (s *Server) Start() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = s.validator

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	e.GET("/health", s.Health)
	e.GET("/docs", s.Docs)

	e.GET("/internal/v1/sessions/:sessionID", s.InternalGetSession)
	e.POST("/internal/v1/sessions", s.InternalCreateSession)
	e.DELETE("/internal/v1/sessions/:sessionID", s.InternalDeleteSession)

	e.POST("/internal/v1/sessions/verify", s.InternalVerifySession)

	e.Any("/extauth", s.ExtAuth)
	e.Any("/extauth/*", s.ExtAuth)

	s.RunWithGracefulShutdown(e)
}

func (s *Server) RunWithGracefulShutdown(e *echo.Echo) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(":" + s.port); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	<-ctx.Done()

	log.Info().Msg("server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server shutdown unexpectedly")
	}
}
