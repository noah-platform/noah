package handler

import (
	"github.com/noah-platform/noah/pkg/validator"

	"github.com/noah-platform/noah/notification/consumer-email/core/port"
)

type Handler struct {
	service   port.Service
	validator *validator.Validator
}

type Dependencies struct {
	Service   port.Service
	Validator *validator.Validator
}

type Config struct {
}

func New(deps Dependencies, cfg Config) *Handler {
	return &Handler{
		service:   deps.Service,
		validator: deps.Validator,
	}
}
