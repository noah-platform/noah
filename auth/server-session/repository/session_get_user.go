package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
)

func (r *SessionRepository) GetUserIDFromSession(ctx context.Context, id string) (string, error) {
	l := log.Ctx(ctx)

	userID, err := r.rdb.Get(ctx, redisSessionKey(id)).Result()
	if err == nil {
		l.Info().Msg("[SessionRepository.GetUserIDFromSession] session loaded from redis")

		if userID == "" {
			l.Info().Msg("[SessionRepository.GetUserIDFromSession] negative cache loaded from redis")

			return "", core.ErrSessionNotFound
		}
		return userID, nil
	}

	if !errors.Is(err, redis.Nil) {
		l.Warn().Err(err).Msg("[SessionRepository.GetUserIDFromSession] failed to get session from redis, trying database")
	}

	session, err := r.queries.GetSession(ctx, id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		l.Error().Err(err).Msg("[SessionRepository.GetUserIDFromSession] failed to get session from database")

		return "", errors.Wrap(err, "failed to get session from database")
	}

	l.Info().Msg("[SessionRepository.GetUserIDFromSession] session loaded from database")

	var ttl time.Duration = 0
	if errors.Is(err, pgx.ErrNoRows) {
		ttl = 10 * time.Minute
	}
	if err := r.rdb.Set(ctx, redisSessionKey(id), session.UserID, ttl).Err(); err != nil {
		l.Warn().Err(err).Msg("[SessionRepository.GetUserIDFromSession] failed to cache session in redis, skipping")
	}

	if errors.Is(err, pgx.ErrNoRows) {
		l.Info().Msg("[SessionRepository.GetUserIDFromSession] session not found")

		return "", core.ErrSessionNotFound
	}
	return session.UserID, nil
}
