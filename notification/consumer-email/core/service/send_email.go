package service

import (
	"context"
	"slices"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/notification/consumer-email/core"
)

func (s *Service) SendEmail(ctx context.Context, msg core.OutgoingEmailMessage) error {
	l := log.Ctx(ctx)

	if !slices.Contains(s.config.AllowedFromAddresses, msg.From) {
		l.Warn().Msg("[Service.SendEmail] from address is not allowed, skipping")

		return nil
	}

	if err := s.mailer.Send(ctx, msg); err != nil {
		l.Error().Err(err).Msg("[Service.SendEmail] failed to send email")

		return errors.Wrap(err, "failed to send email")
	}

	l.Info().Msg("[Service.SendEmail] mail sent")

	return nil
}
