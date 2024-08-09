package handler

import (
	"github.com/noah-platform/noah/account/server/core/port"
	"github.com/noah-platform/noah/pkg/validator"
)

type Server struct {
	service   port.Service
	validator *validator.Validator

	port      string
	jwtSecret string
}

type Dependencies struct {
	Service   port.Service
	Validator *validator.Validator
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
