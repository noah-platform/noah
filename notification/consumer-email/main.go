package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/logging"

	"github.com/noah-platform/noah/notification/consumer-email/di"
)

type Config struct {
	Environment string `env:"APP_ENV,required"`

	KafkaBrokers []string `env:"KAFKA_BROKERS,required"`
	KafkaTopics  []string `env:"KAFKA_TOPICS,required,"`
	KafkaGroupID string   `env:"KAFKA_GROUP_ID,required"`
}

func init() {
	logging.Init()
}

func main() {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		log.Panic().Err(err).Msg("failed to load config")
	}

	log.Info().Str("environment", cfg.Environment).Msg("consumer is starting")

	consumer := di.New(di.Config{
		ConsumerConfig: di.ConsumerConfig{
			ClientID: "notification-consumer-email",
			Brokers:  cfg.KafkaBrokers,
			Topics:   cfg.KafkaTopics,
			GroupID:  cfg.KafkaGroupID,
		},
		ServiceConfig: di.ServiceConfig{},
	})

	consumer.Start()
}
