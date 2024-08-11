package service

import "github.com/noah-platform/noah/notification/consumer-email/core/port"

type Service struct {
	config               Config
	mailer               port.Mailer
	allowedFromAddresses []string
}

type Config struct {
	AllowedFromAddresses []string
}

type Dependencies struct {
	Mailer port.Mailer
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config:               cfg,
		mailer:               deps.Mailer,
		allowedFromAddresses: cfg.AllowedFromAddresses,
	}
}
