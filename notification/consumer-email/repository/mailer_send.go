package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/wneessen/go-mail"

	"github.com/noah-platform/noah/notification/consumer-email/core"
)

func (m *Mailer) Send(ctx context.Context, email core.OutgoingEmailMessage) error {
	l := log.Ctx(ctx)

	message := mail.NewMsg()
	message.SetMessageID()
	messageID := message.GetGenHeader(mail.HeaderMessageID)

	if err := message.FromFormat(email.SenderName, email.From); err != nil {
		log.Error().Err(err).Msg("[MailerRepository.Send] failed to set From address")

		return errors.Wrap(err, "failed to set From address")
	}
	if err := message.AddToFormat(email.RecipientName, email.To); err != nil {
		log.Error().Err(err).Msg("[MailerRepository.Send] failed to set To address")

		return errors.Wrap(err, "failed to set To address")
	}
	message.Subject(email.Subject)
	message.SetBodyString(mail.TypeTextHTML, email.Body)

	if err := m.mailer.DialAndSendWithContext(ctx, message); err != nil {
		log.Error().Err(err).Msg("[MailerRepository.Send] failed to send mail")

		return errors.Wrap(err, "failed to send mail")
	}

	l.Info().Strs("messageId", messageID).Msg("[MailerRepository.Send] mail sent")

	return nil
}
