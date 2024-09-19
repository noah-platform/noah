package service

import "github.com/noah-platform/noah/auth/server-session/core/port"

type Service struct {
	config      Config
	sessionRepo port.SessionRepository
}

type Config struct {
	SessionIDLength int
}

type Dependencies struct {
	SessionRepository port.SessionRepository
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config:      cfg,
		sessionRepo: deps.SessionRepository,
	}
}
