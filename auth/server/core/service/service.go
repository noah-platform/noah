package service

import (
	accountClient "github.com/noah-platform/noah/account/server/client"
	authSessionClient "github.com/noah-platform/noah/auth/server-session/client"
	"google.golang.org/api/idtoken"
)

type Service struct {
	config                 Config
	accountClient          *accountClient.Client
	authSessionClient      *authSessionClient.Client
	googleIDTokenValidator *idtoken.Validator
}

type Config struct {
	GoogleClientID string
}

type Dependencies struct {
	AccountClient          *accountClient.Client
	AuthSessionClient      *authSessionClient.Client
	GoogleIDTokenValidator *idtoken.Validator
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config:                 cfg,
		accountClient:          deps.AccountClient,
		authSessionClient:      deps.AuthSessionClient,
		googleIDTokenValidator: deps.GoogleIDTokenValidator,
	}
}
