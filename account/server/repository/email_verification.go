package repository

import (
	"context"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/messaging"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (e *EmailRepository) ProduceOutgoingEmailVerificationMessage(ctx context.Context, traceID string, message core.OutgoingEmailMessage) error {
	l := log.Ctx(ctx)

	partition, offset, err := e.producer.SendMessage(e.topic, messaging.EventOutgoingEmail, traceID, message)
	if err != nil {
		l.Error().Err(err).Msg("[EmailRepository.ProduceEmailVerificationRequest] failed to produce message")

		return errors.Wrap(err, "failed to produce message")
	}

	l.Info().Int("partition", partition).Int("offset", offset).Msg("[EmailRepository.ProduceEmailVerificationRequest] produced email verification request")

	return nil
}
