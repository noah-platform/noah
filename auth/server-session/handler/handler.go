package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/noah-platform/noah/auth/server-session/core/port"
)

type Server struct {
	service   port.Service
	validator echo.Validator

	port         string
	jwtSecret    string
	cookieDomain string
}

type Dependencies struct {
	Service   port.Service
	Validator echo.Validator
}

type Config struct {
	Port         string
	JWTSecret    string
	CookieDomain string
}

func New(deps Dependencies, cfg Config) *Server {
	return &Server{
		service:   deps.Service,
		validator: deps.Validator,

		port:         cfg.Port,
		jwtSecret:    cfg.JWTSecret,
		cookieDomain: cfg.CookieDomain,
	}
}
