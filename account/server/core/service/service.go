package service

import "github.com/noah-platform/noah/account/server/core/port"

type Service struct {
	config      Config
	accountRepo port.AccountRepository
	emailRepo   port.EmailRepository
}

type Config struct {
	EmailFrom string
}

type Dependencies struct {
	AccountRepository port.AccountRepository
	EmailRepository   port.EmailRepository
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config:      cfg,
		accountRepo: deps.AccountRepository,
		emailRepo:   deps.EmailRepository,
	}
}
