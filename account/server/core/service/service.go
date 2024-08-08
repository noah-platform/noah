package service

import "github.com/noah-platform/noah/account/server/core/port"

type Service struct {
	config      Config
	accountRepo port.AccountRepository
}

type Config struct {
}

type Dependencies struct {
	AccountRepository port.AccountRepository
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config:      cfg,
		accountRepo: deps.AccountRepository,
	}
}
