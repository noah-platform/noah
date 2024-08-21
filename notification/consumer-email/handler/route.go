package handler

import (
	"context"
	"encoding/json"

	"github.com/noah-platform/noah/pkg/messaging"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Handle(ctx context.Context, event messaging.Event, payload json.RawMessage) error {
	l := log.With().Str("traceId", ctx.Value(messaging.TraceIDContextKey).(string)).Logger()
	ctx = l.WithContext(ctx)

	switch event {
	case messaging.EventOutgoingEmail:
		l.Info().Msg("[Handler.Handle] received outgoing email event")

		return h.OutgoingEmail(ctx, payload)
	default:
		l.Warn().Msg("[Handler.Handle] received unknown event, skipping")

		return nil
	}
}
