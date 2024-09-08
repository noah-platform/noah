package service

import (
	accountClient "github.com/noah-platform/noah/account/server/client"
	"google.golang.org/api/idtoken"
)

type Service struct {
	config                 Config
	accountClient          *accountClient.Client
	googleIDTokenValidator *idtoken.Validator
}

type Config struct {
}

type Dependencies struct {
	AccountClient          *accountClient.Client
	GoogleIDTokenValidator *idtoken.Validator
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config:                 cfg,
		accountClient:          deps.AccountClient,
		googleIDTokenValidator: deps.GoogleIDTokenValidator,
	}
}
