package service

import (
	accountClient "github.com/noah-platform/noah/account/server/client"
)

type Service struct {
	config        Config
	accountClient *accountClient.Client
}

type Config struct {
}

type Dependencies struct {
	AccountClient *accountClient.Client
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config:        cfg,
		accountClient: deps.AccountClient,
	}
}
