package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/di"
	"github.com/noah-platform/noah/pkg/logging"
)

type Config struct {
	Environment string `env:"APP_ENV,required"`

	Port         string   `env:"PORT,required"`
	JWTSecret    string   `env:"JWT_SECRET,required"`
	DatabaseUrl  string   `env:"DATABASE_URL,required"`
	KafkaBrokers []string `env:"KAFKA_BROKERS,required"`

	EmailFrom       string `env:"EMAIL_FROM,required"`
	EmailKafkaTopic string `env:"EMAIL_KAFKA_TOPIC,required"`
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
		ServiceConfig: di.ServiceConfig{
			EmailFrom: cfg.EmailFrom,
		},
		PostgresConfig: di.PostgresConfig{
			DatabaseUrl: cfg.DatabaseUrl,
		},
		ProducerConfig: di.ProducerConfig{
			Brokers:  cfg.KafkaBrokers,
			ClientID: "account-server",
		},
		EmailRepoConfig: di.EmailRepoConfig{
			KafkaTopic: cfg.EmailKafkaTopic,
		},
	})

	server.Start()
}
