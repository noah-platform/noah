package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/logging"

	"github.com/noah-platform/noah/auth/server/di"
)

type Config struct {
	Environment string `env:"APP_ENV,required"`

	Port      string `env:"PORT,required"`
	JWTSecret string `env:"JWT_SECRET,required"`

	AccountServerURL string `env:"ACCOUNT_SERVER_URL,required"`
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
		AccountClientConfig: di.AccountClientConfig{
			BaseURL:  cfg.AccountServerURL,
			RetryMax: 3,
		},
	})

	server.Start()
}
