package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
)

func (r *SessionRepository) DeleteSession(ctx context.Context, id string) (err error) {
	l := log.Ctx(ctx)

	tx, err := r.db.Begin(ctx)
	if err != nil {
		l.Error().Err(err).Msg("[SessionRepository.DeleteSession] failed to begin transaction")

		return errors.Wrap(err, "failed to begin transaction")
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(ctx); err == nil {
				l.Warn().Msg("[SessionRepository.DeleteSession] transaction rolled back")
			} else {
				l.Error().Err(err).Msg("[SessionRepository.DeleteSession] failed to rollback transaction")
			}
		}
	}()

	_, err = r.queries.WithTx(tx).DeleteSession(ctx, id)
	if err != nil {
		switch {
		case errors.Is(pgx.ErrNoRows, err):
			l.Info().Msg("[SessionRepository.DeleteSession] session not found")

			return core.ErrSessionNotFound
		default:
			l.Error().Err(err).Msg("[SessionRepository.DeleteSession] failed to delete session from database")

			return errors.Wrap(err, "failed to delete session from database")
		}
	}

	if err := r.rdb.Del(ctx, redisSessionKey(id)).Err(); err != nil {
		l.Error().Err(err).Msg("[SessionRepository.DeleteSession] failed to delete session from redis")

		return errors.Wrap(err, "failed to delete session from redis")
	}

	if err := tx.Commit(ctx); err != nil {
		l.Error().Err(err).Msg("[SessionRepository.DeleteSession] failed to commit transaction")

		return errors.Wrap(err, "failed to commit transaction")
	}

	l.Info().Msg("[SessionRepository.DeleteSession] session deleted")

	return nil
}
