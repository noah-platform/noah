package di

import (
	"github.com/noah-platform/noah/pkg/producer"
	"github.com/noah-platform/noah/pkg/validator"

	"github.com/noah-platform/noah/account/server/core/service"
	"github.com/noah-platform/noah/account/server/handler"
	"github.com/noah-platform/noah/account/server/repository"
)

type ServerConfig = handler.Config

type ServiceConfig = service.Config

type EmailRepoConfig = repository.EmailRepoConfig

type Config struct {
	ServerConfig    ServerConfig
	ServiceConfig   ServiceConfig
	PostgresConfig  PostgresConfig
	ProducerConfig  ProducerConfig
	EmailRepoConfig EmailRepoConfig
}

func New(cfg Config) *handler.Server {
	validator := validator.NewValidator()

	producer := newProducer(producer.Dependencies{
		Validator: validator,
	}, cfg.ProducerConfig)

	pgx := newPgxClient(cfg.PostgresConfig)
	accountRepo := repository.NewAccountRepository(repository.AccountRepoDependencies{
		PgxClient: pgx,
	})

	emailRepo := repository.NewEmailRepository(repository.EmailRepoDependencies{
		Producer: producer,
	}, cfg.EmailRepoConfig)

	service := service.New(service.Dependencies{
		AccountRepository: accountRepo,
		EmailRepository:   emailRepo,
	}, cfg.ServiceConfig)

	server := handler.New(handler.Dependencies{
		Service:   service,
		Validator: validator,
	}, cfg.ServerConfig)

	return server
}
