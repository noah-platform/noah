package handler

import (
	"context"
	"encoding/json"

	"github.com/noah-platform/noah/notification/consumer-email/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (h *Handler) OutgoingEmail(ctx context.Context, payload json.RawMessage) error {
	l := log.Ctx(ctx)

	var email core.OutgoingEmailMessage
	if err := json.Unmarshal(payload, &email); err != nil {
		l.Warn().Err(err).Msg("[Handler.OutgoingEmail] failed to unmarshal message")

		return errors.Wrap(err, "failed to unmarshal message")
	}

	if err := h.validator.Validate(email); err != nil {
		l.Error().Err(err).Msg("[Handler.OutgoingEmail] failed to validate message")

		return errors.Wrap(err, "failed to validate message")
	}

	*l = l.With().Str("from", email.From).Str("to", email.To).Str("subject", email.Subject).Logger()
	ctx = l.WithContext(ctx)

	if err := h.service.SendEmail(ctx, email); err != nil {
		l.Error().Err(err).Msg("[Handler.OutgoingEmail] failed to send email")

		return errors.Wrap(err, "failed to send email")
	}

	l.Info().Msg("[Handler.OutgoingEmail] outgoing email processed")

	return nil
}
