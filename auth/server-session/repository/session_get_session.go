package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
)

func (r *SessionRepository) GetSession(ctx context.Context, id string) (*core.Session, error) {
	l := log.Ctx(ctx)

	session, err := r.queries.GetSession(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			l.Info().Msg("[SessionRepository.GetSession] session not found")

			return nil, core.ErrSessionNotFound
		default:
			l.Error().Err(err).Msg("[SessionRepository.GetSession] failed to get session")

			return nil, errors.Wrap(err, "failed to get session")
		}
	}

	l.Debug().Msg("[SessionRepository.GetSession] session loaded")

	return &core.Session{
		SessionID: session.SessionID,
		UserID:    session.UserID,
		IPAddress: session.IpAddress,
		UserAgent: session.UserAgent,
		CreatedAt: session.CreatedAt.Time,
	}, nil
}
