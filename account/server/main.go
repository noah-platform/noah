package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/di"
)

type Config struct {
	Environment string `env:"APP_ENV,required"`

	Port        string `env:"PORT,required"`
	JWTSecret   string `env:"JWT_SECRET,required"`
	DatabaseUrl string `env:"DATABASE_URL,required"`
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
			DatabaseUrl: cfg.DatabaseUrl,
		},
	})

	server.Start()
}
