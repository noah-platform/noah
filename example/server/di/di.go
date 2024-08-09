package di

import (
	"github.com/noah-platform/noah/pkg/validator"

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
	MongoConfig    MongoConfig
}

func New(cfg Config) *handler.Server {
	// pgx := newPgxClient(cfg.PostgresConfig)
	// examplePostgresRepo = repository.NewExamplePostgresRepository(repository.ExamplePostgresRepoDependencies{
	// 	PgxClient: pgx,
	// })

	db := newMongoClient(cfg.MongoConfig)
	exampleMongoRepo := repository.NewExampleMongoRepository(repository.ExampleMongoRepoDependencies{
		Example: db.Collection("example"),
	})

	service := service.New(service.Dependencies{
		ExampleRepository: exampleMongoRepo,
	}, cfg.ServiceConfig)

	server := handler.New(handler.Dependencies{
		Service:   service,
		Validator: validator.NewValidator(),
	}, cfg.ServerConfig)

	return server
}
