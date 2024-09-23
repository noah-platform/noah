package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/noah-platform/noah/auth/server/core/port"
)

const sessionCookieName = "noahses"

type Server struct {
	service   port.Service
	validator echo.Validator

	port      string
	jwtSecret string
}

type Dependencies struct {
	Service   port.Service
	Validator echo.Validator
}

type Config struct {
	Port      string
	JWTSecret string
}

func New(deps Dependencies, cfg Config) *Server {
	return &Server{
		service:   deps.Service,
		validator: deps.Validator,

		port:      cfg.Port,
		jwtSecret: cfg.JWTSecret,
	}
}
