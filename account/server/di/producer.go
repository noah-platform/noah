package di

import (
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/producer"
)

type ProducerDependencies = producer.Dependencies

type ProducerConfig = producer.Config

func newProducer(deps ProducerDependencies, cfg ProducerConfig) *producer.Producer {
	producer, err := producer.NewProducer(deps, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create kafka producer")
	}

	return producer
}
