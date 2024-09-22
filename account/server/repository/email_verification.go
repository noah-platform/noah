package repository

import (
	"context"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/messaging"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (e *EmailRepository) ProduceOutgoingEmail(ctx context.Context, traceID string, message core.OutgoingEmailMessage) error {
	l := log.Ctx(ctx)

	partition, offset, err := e.producer.SendMessage(e.topic, messaging.EventOutgoingEmail, traceID, message)
	if err != nil {
		l.Error().Err(err).Msg("[EmailRepository.ProduceOutgoingEmail] failed to produce outgoing email")

		return errors.Wrap(err, "failed to produce outgoing email")
	}

	l.Info().Int("partition", partition).Int("offset", offset).Msg("[EmailRepository.ProduceOutgoingEmail] produced outgoing email")

	return nil
}
