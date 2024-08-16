package producer

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/messaging"
	"github.com/noah-platform/noah/pkg/validator"
)

type Producer struct {
	producer  sarama.SyncProducer
	validator *validator.Validator
}

type Dependencies struct {
	Validator *validator.Validator
}

type Config struct {
	Brokers  []string
	ClientID string
}

func NewProducer(deps Dependencies, cfg Config) (*Producer, error) {
	config := sarama.NewConfig()
	config.ClientID = cfg.ClientID
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(cfg.Brokers, config)
	if err != nil {
		log.Error().Strs("brokers", cfg.Brokers).Str("clientId", cfg.ClientID).Err(err).Msg("failed to initialize kafka producer")

		return nil, errors.Wrap(err, "failed to initialize kafka producer")
	}

	return &Producer{
		producer:  producer,
		validator: deps.Validator,
	}, nil
}

func (p *Producer) SendMessage(topic string, event messaging.Event, traceID string, payload any) (int, int, error) {
	l := log.With().Str("topic", topic).Str("event", string(event)).Str("traceID", traceID).Logger()

	message := messaging.ProducerMessage{
		TraceID:   traceID,
		Event:     event,
		Payload:   payload,
		Timestamp: time.Now(),
	}
	if err := p.validator.Validate(message); err != nil {
		l.Error().Err(err).Msg("failed to validate message")

		return 0, 0, errors.Wrap(err, "failed to validate message")
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		l.Error().Err(err).Msg("failed to marshal message")

		return 0, 0, errors.Wrap(err, "failed to marshal message")
	}

	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(jsonMessage),
	}
	partition, offset, err := p.producer.SendMessage(kafkaMessage)
	if err != nil {
		l.Error().Err(err).Msg("failed to produce message")

		return 0, 0, errors.Wrap(err, "failed to produce message")
	}

	return int(partition), int(offset), nil
}
