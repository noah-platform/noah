package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
	"github.com/noah-platform/noah/auth/server-session/generated/sqlc"
)

func (r *SessionRepository) CreateSession(ctx context.Context, input *core.Session) (*core.Session, error) {
	l := log.Ctx(ctx)

	params := sqlc.CreateSessionParams{
		SessionID: input.SessionID,
		UserID:    input.UserID,
		IpAddress: input.IPAddress,
		UserAgent: input.UserAgent,
	}
	session, err := r.queries.CreateSession(ctx, params)
	if err != nil {
		l.Error().Err(err).Msg("[SessionRepository.CreateSession] failed to create session")

		return nil, errors.Wrap(err, "failed to create session")
	}

	if err := r.rdb.Set(ctx, redisSessionKey(session.SessionID), session.UserID, 0).Err(); err != nil {
		l.Warn().Err(err).Msg("[SessionRepository.CreateSession] cannot store session in redis, skipping")
	}

	l.Info().Msg("[SessionRepository.CreateSession] session created")

	return &core.Session{
		SessionID: session.SessionID,
		UserID:    session.UserID,
		IPAddress: session.IpAddress,
		UserAgent: session.UserAgent,
		CreatedAt: session.CreatedAt.Time,
	}, nil
}
