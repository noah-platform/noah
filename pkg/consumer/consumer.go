package consumer

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/messaging"
	"github.com/noah-platform/noah/pkg/validator"
)

type Consumer struct {
	config *sarama.Config

	brokers []string
	topics  []string
	groupID string

	handler   Handler
	validator *validator.Validator

	ready chan bool
}

type Dependencies struct {
	Handler   Handler
	Validator *validator.Validator
}

type Config struct {
	ClientID string
	Brokers  []string
	Topics   []string
	GroupID  string
}

func NewConsumer(deps Dependencies, cfg Config) *Consumer {
	config := sarama.NewConfig()
	config.ClientID = cfg.ClientID
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	return &Consumer{
		config: config,

		brokers: cfg.Brokers,
		topics:  cfg.Topics,
		groupID: cfg.GroupID,

		handler:   deps.Handler,
		validator: deps.Validator,

		ready: make(chan bool),
	}
}

func (c *Consumer) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	l := log.With().Strs("brokers", c.brokers).Strs("topics", c.topics).Str("groupId", c.groupID).Logger()

	client, err := sarama.NewConsumerGroup(c.brokers, c.groupID, c.config)
	if err != nil {
		l.Fatal().Err(err).Msg("failed to create consumer group")
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, c.topics, c); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}

				l.Fatal().Err(err).Msg("failed to join consumer group")
			}

			if ctx.Err() != nil {
				return
			}

			c.ready = make(chan bool)
		}
	}()

	<-c.ready

	l.Info().Msg("consumer is up and running")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	keepRunning := true
	for keepRunning {
		select {
		case <-ctx.Done():
			l.Info().Msg("consumer is shutting down due to context cancellation")

			keepRunning = false
		case <-sigterm:
			l.Info().Msg("consumer is shutting down")

			keepRunning = false
		}
	}
	cancel()

	wg.Wait()

	if err = client.Close(); err != nil {
		l.Fatal().Err(err).Msg("failed to close consumer")
	}
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Info().Msg("message channel was closed")

				return nil
			}

			var event messaging.Message
			if err := json.Unmarshal(message.Value, &event); err != nil {
				log.Fatal().Err(err).Str("topic", message.Topic).Int32("partition", message.Partition).Int64("offset", message.Offset).Msg("failed to unmarshal message")

				return errors.Wrap(err, "failed to unmarshal message")
			}

			l := log.With().
				Str("traceId", event.TraceID).
				Str("event", string(event.Event)).
				Str("topic", message.Topic).
				Int32("partition", message.Partition).
				Int64("offset", message.Offset).
				Time("timestamp", event.Timestamp).
				Logger()
			l.Info().Msg("received message")

			if err := c.validator.Validate(event); err != nil {
				l.Fatal().Err(err).Msg("failed to validate message")

				return errors.Wrap(err, "failed to validate message")
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, messaging.TraceIDContextKey, event.TraceID)
			if err := c.handler.Handle(ctx, event.Event, event.Payload); err != nil {
				l.Fatal().Err(err).Msg("failed to handle message")

				return errors.Wrap(err, "failed to handle message")
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
