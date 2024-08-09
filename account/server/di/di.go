package di

import (
	"github.com/noah-platform/noah/pkg/validator"

	"github.com/noah-platform/noah/account/server/core/service"
	"github.com/noah-platform/noah/account/server/handler"
	"github.com/noah-platform/noah/account/server/repository"
)

type ServerConfig = handler.Config

type ServiceConfig = service.Config

type Config struct {
	ServerConfig   ServerConfig
	ServiceConfig  ServiceConfig
	PostgresConfig PostgresConfig
}

func New(cfg Config) *handler.Server {
	pgx := newPgxClient(cfg.PostgresConfig)
	accountRepo := repository.NewAccountRepository(repository.AccountRepoDependencies{
		PgxClient: pgx,
	})

	service := service.New(service.Dependencies{
		AccountRepository: accountRepo,
	}, cfg.ServiceConfig)

	server := handler.New(handler.Dependencies{
		Service:   service,
		Validator: validator.NewValidator(),
	}, cfg.ServerConfig)

	return server
}
