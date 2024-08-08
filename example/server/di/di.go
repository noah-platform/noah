package di

import (
	"github.com/noah-platform/noah/example/server/core/service"
	"github.com/noah-platform/noah/example/server/handler"
	"github.com/noah-platform/noah/example/server/repository"
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
	exampleRepo := repository.NewExamplePostgresRepository(repository.ExamplePostgresRepoDependencies{
		PgxClient: pgx,
	})

	service := service.New(service.Dependencies{
		ExampleRepository: exampleRepo,
	}, cfg.ServiceConfig)

	server := handler.New(handler.Dependencies{
		Service: service,
	}, cfg.ServerConfig)

	return server
}
