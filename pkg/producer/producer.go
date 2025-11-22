package producer

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"github.com/riferrei/srclient"
	"github.com/rs/zerolog/log"
	"github.com/xdg-go/scram"

	"github.com/noah-platform/noah/pkg/messaging"
	"github.com/noah-platform/noah/pkg/validator"
)

type Producer struct {
	producer  sarama.SyncProducer
	srclient  srclient.ISchemaRegistryClient
	validator *validator.Validator
}

type Dependencies struct {
	Validator *validator.Validator
}

type Config struct {
	Brokers           []string
	SchemaRegistryURL string
	ClientID          string
	EnforceTLS        bool
	AuthEnabled       bool
	Username          string
	Password          string
}

func NewProducer(deps Dependencies, cfg Config) (*Producer, error) {
	config := sarama.NewConfig()
	config.ClientID = cfg.ClientID
	config.Producer.Return.Successes = true
	if cfg.EnforceTLS {
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: false,
		}
	}
	if cfg.AuthEnabled {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = cfg.Username
		config.Net.SASL.Password = cfg.Password
		config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
		config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &XDGSCRAMClient{HashGeneratorFcn: scram.SHA512}
		}
	}

	producer, err := sarama.NewSyncProducer(cfg.Brokers, config)
	if err != nil {
		log.Error().Strs("brokers", cfg.Brokers).Str("clientId", cfg.ClientID).Err(err).Msg("failed to initialize kafka producer")

		return nil, errors.Wrap(err, "failed to initialize kafka producer")
	}

	schemaRegistryClient := srclient.CreateSchemaRegistryClient(cfg.SchemaRegistryURL)

	return &Producer{
		producer:  producer,
		srclient:  schemaRegistryClient,
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

func (p *Producer) SendAvroMessage(topic string, traceID string, message messaging.AvroMessage) (int, int, error) {
	l := log.With().Str("topic", topic).Logger()

	schema, err := p.getSchema(topic, message.Schema())
	if err != nil {
		l.Error().Err(err).Msg("failed to get or register schema")

		return 0, 0, errors.Wrap(err, "failed to get or register schema")
	}

	value, err := p.serializeAvroMessage(schema.ID(), message)
	if err != nil {
		l.Error().Err(err).Msg("failed to serialize avro message")

		return 0, 0, errors.Wrap(err, "failed to serialize avro message")
	}

	producerMessage := &sarama.ProducerMessage{
		Topic: topic,
		Headers: []sarama.RecordHeader{
			{Key: []byte("traceId"), Value: []byte(traceID)},
			{Key: []byte("timestamp"), Value: []byte(time.Now().UTC().Format(time.RFC3339Nano))},
		},
		Value: sarama.ByteEncoder(value),
	}
	partition, offset, err := p.producer.SendMessage(producerMessage)
	if err != nil {
		l.Error().Err(err).Msg("failed to produce avro message")

		return 0, 0, errors.Wrap(err, "failed to produce avro message")
	}

	return int(partition), int(offset), nil
}

func (p *Producer) getSchema(topic string, schemaStr string) (*srclient.Schema, error) {
	schema, err := p.srclient.GetLatestSchema(topic)
	if err != nil {
		if e, ok := err.(srclient.Error); !ok || e.Code != 40401 {
			log.Error().Err(err).Msg("failed to get latest schema from schema registry")

			return nil, errors.Wrap(err, "failed to get latest schema from schema registry")
		}
	}

	// TODO: Register new schema on change
	if schema == nil {
		schema, err = p.srclient.CreateSchema(topic, schemaStr, srclient.Avro)
		if err != nil {
			log.Error().Err(err).Msg("failed to register schema in schema registry")

			return nil, errors.Wrap(err, "failed to register schema in schema registry")
		}

		log.Info().Str("topic", topic).Int("schemaId", schema.ID()).Msg("registered new schema in schema registry")
	}

	return schema, nil
}

func (p *Producer) serializeAvroMessage(schemaID int, message messaging.AvroMessage) ([]byte, error) {
	var buf bytes.Buffer

	// Magic Byte
	buf.WriteByte(0)

	// Schema ID
	err := binary.Write(&buf, binary.BigEndian, uint32(schemaID))
	if err != nil {
		log.Error().Err(err).Msg("failed to write schema ID to buffer")

		return nil, errors.Wrap(err, "failed to write schema ID to buffer")
	}

	// Serialized Payload
	var payload bytes.Buffer
	if err := message.Serialize(&payload); err != nil {
		log.Error().Err(err).Msg("failed to serialize avro message")

		return nil, errors.Wrap(err, "failed to serialize avro message")
	}
	buf.Write(payload.Bytes())

	return buf.Bytes(), nil
}
