package repository

import (
	"context"

	"github.com/noah-platform/noah/pkg/messaging"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (e *EmailRepository) ProduceEmailVerificationRequest(ctx context.Context, to, name, url string) error {
	l := log.Ctx(ctx)

	payload := messaging.OutgoingEmailMessage{
		From:          e.emailFrom,
		SenderName:    "Noah Platform",
		To:            to,
		RecipientName: name,
		Subject:       "Verify your email",
		Body:          "Hello " + name + ",\n\n" + "Please verify your email address by clicking the following link: " + url + "\n\n" + "Thank you,\n" + "Noah Platform", // TODO: Use a template
	}
	partition, offset, err := e.producer.SendMessage(e.topic, messaging.EventOutgoingEmail, payload, "TODO")
	if err != nil {
		l.Error().Err(err).Msg("[EmailRepository.ProduceEmailVerificationRequest] failed to produce message")

		return errors.Wrap(err, "failed to produce message")
	}

	l.Info().Int("partition", partition).Int("offset", offset).Msg("[EmailRepository.ProduceEmailVerificationRequest] produced email verification request")

	return nil
}
