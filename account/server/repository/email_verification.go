package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/account/server/generated/avro"
	"github.com/noah-platform/noah/pkg/messaging"
)

func (e *EmailRepository) ProduceOutgoingEmail(ctx context.Context, traceID string, message core.OutgoingEmailMessage) error {
	l := log.Ctx(ctx)

	producerMessage := avro.OutgoingEmailEvent{
		Event: string(messaging.EventOutgoingEmail),
		Payload: avro.EmailPayload{
			To:            message.To,
			SenderName:    message.SenderName,
			From:          message.From,
			RecipientName: message.RecipientName,
			Subject:       message.Subject,
			Body:          message.Body,
		},
	}
	partition, offset, err := e.producer.SendAvroMessage(e.topic, traceID, producerMessage)
	if err != nil {
		l.Error().Err(err).Msg("[EmailRepository.ProduceOutgoingEmail] failed to produce outgoing email")

		return errors.Wrap(err, "failed to produce outgoing email")
	}

	l.Info().Int("partition", partition).Int("offset", offset).Msg("[EmailRepository.ProduceOutgoingEmail] produced outgoing email")

	return nil
}
