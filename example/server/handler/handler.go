package handler

import (
	"github.com/go-playground/validator/v10"

	"github.com/noah-platform/noah/example/server/core/port"
)

type Server struct {
	service   port.Service
	validator *validator.Validate

	port      string
	jwtSecret string
}

type Dependencies struct {
	Service port.Service
}

type Config struct {
	Port      string
	JWTSecret string
}

func New(deps Dependencies, cfg Config) *Server {
	return &Server{
		service:   deps.Service,
		validator: validator.New(),

		port:      cfg.Port,
		jwtSecret: cfg.JWTSecret,
	}
}
