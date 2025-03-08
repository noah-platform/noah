package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/noah-platform/noah/question-bank/server/core/port"
)

type Server struct {
	service   port.Service
	validator echo.Validator
	auth      AuthMiddleware

	port      string
	jwtSecret string
}

type Dependencies struct {
	Service        port.Service
	Validator      echo.Validator
	AuthMiddleware AuthMiddleware
}

type Config struct {
	Port      string
	JWTSecret string
}

func New(deps Dependencies, cfg Config) *Server {
	return &Server{
		service:   deps.Service,
		validator: deps.Validator,
		auth:      deps.AuthMiddleware,

		port:      cfg.Port,
		jwtSecret: cfg.JWTSecret,
	}
}
