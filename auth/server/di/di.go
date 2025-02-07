package di

import (
	"github.com/noah-platform/noah/pkg/validator"

	accountClient "github.com/noah-platform/noah/account/server/client"
	authSessionClient "github.com/noah-platform/noah/auth/server-session/client"

	"github.com/noah-platform/noah/auth/server/core/service"
	"github.com/noah-platform/noah/auth/server/handler"
)

type ServerConfig = handler.Config

type ServiceConfig = service.Config

type AccountClientConfig = accountClient.Config

type AuthSessionClientConfig = authSessionClient.Config

type Config struct {
	ServerConfig            ServerConfig
	ServiceConfig           ServiceConfig
	AccountClientConfig     AccountClientConfig
	AuthSessionClientConfig AuthSessionClientConfig
}

func New(cfg Config) *handler.Server {
	accountClient := accountClient.New(cfg.AccountClientConfig)
	authSessionClient := authSessionClient.New(cfg.AuthSessionClientConfig)
	googleIDTokenValidator := newGoogleIDTokenValidator()
	service := service.New(service.Dependencies{
		AccountClient:          accountClient,
		AuthSessionClient:      authSessionClient,
		GoogleIDTokenValidator: googleIDTokenValidator,
	}, cfg.ServiceConfig)

	server := handler.New(handler.Dependencies{
		Service:   service,
		Validator: validator.NewValidator(),
	}, cfg.ServerConfig)

	return server
}
