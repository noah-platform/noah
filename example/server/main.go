package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/example/server/di"
	"github.com/noah-platform/noah/pkg/logging"
)

type Config struct {
	Environment string `env:"APP_ENV,required"`

	Port        string `env:"PORT,required"`
	JWTSecret   string `env:"JWT_SECRET,required"`
	PostgresUrl string `env:"POSTGRES_URL,required"`
	MongoUri    string `env:"MONGO_URI,required"`
}

func init() {
	logging.Init()
}

func main() {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		log.Panic().Err(err).Msg("failed to load config")
	}

	log.Info().Str("environment", cfg.Environment).Msg("server is starting")

	server := di.New(di.Config{
		ServerConfig: di.ServerConfig{
			Port:      cfg.Port,
			JWTSecret: cfg.JWTSecret,
		},
		ServiceConfig: di.ServiceConfig{},
		PostgresConfig: di.PostgresConfig{
			DatabaseUrl: cfg.PostgresUrl,
		},
		MongoConfig: di.MongoConfig{
			DatabaseUri: cfg.MongoUri,
		},
	})

	server.Start()
}
