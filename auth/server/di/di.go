package di

import (
	"github.com/noah-platform/noah/pkg/validator"

	accountClient "github.com/noah-platform/noah/account/server/client"

	"github.com/noah-platform/noah/auth/server/core/service"
	"github.com/noah-platform/noah/auth/server/handler"
)

type ServerConfig = handler.Config

type ServiceConfig = service.Config

type AccountClientConfig = accountClient.Config

type Config struct {
	ServerConfig        ServerConfig
	ServiceConfig       ServiceConfig
	AccountClientConfig AccountClientConfig
}

func New(cfg Config) *handler.Server {
	accountClient := accountClient.New(cfg.AccountClientConfig)
	service := service.New(service.Dependencies{
		AccountClient: accountClient,
	}, cfg.ServiceConfig)

	server := handler.New(handler.Dependencies{
		Service:   service,
		Validator: validator.NewValidator(),
	}, cfg.ServerConfig)

	return server
}
