package di

import (
	"github.com/noah-platform/noah/pkg/validator"

	"github.com/noah-platform/noah/auth/server-session/core/service"
	"github.com/noah-platform/noah/auth/server-session/handler"
	"github.com/noah-platform/noah/auth/server-session/repository"
)

type ServerConfig = handler.Config

type ServiceConfig = service.Config

type Config struct {
	ServerConfig   ServerConfig
	ServiceConfig  ServiceConfig
	PostgresConfig PostgresConfig
	RedisConfig    RedisConfig
}

func New(cfg Config) *handler.Server {
	pgx := newPgxClient(cfg.PostgresConfig)
	rdb := newRedisClient(cfg.RedisConfig)
	sessionRepo := repository.NewSessionRepository(repository.SessionRepoDependencies{
		PgxClient:   pgx,
		RedisClient: rdb,
	})

	service := service.New(service.Dependencies{
		SessionRepository: sessionRepo,
	}, cfg.ServiceConfig)

	server := handler.New(handler.Dependencies{
		Service:   service,
		Validator: validator.NewValidator(),
	}, cfg.ServerConfig)

	return server
}
