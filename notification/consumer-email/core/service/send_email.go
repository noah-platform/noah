package service

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/notification/consumer-email/core"
)

func (s *Service) SendEmail(ctx context.Context, msg core.OutgoingEmailMessage) error {
	l := log.Ctx(ctx)

	// TODO: Implement email sending logic

	l.Info().Msg("[Service.SendEmail] email sent")

	return nil
}
