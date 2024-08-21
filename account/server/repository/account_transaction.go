package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *AccountRepository) BeginTransaction(ctx context.Context) (pgx.Tx, error) {
	l := log.Ctx(ctx)

	tx, err := r.db.Begin(ctx)
	if err != nil {
		l.Error().Err(err).Msg("[AccountRepository.BeginTransaction] failed to begin transaction")

		return nil, errors.Wrap(err, "failed to begin transaction")
	}

	return tx, nil
}

func (r *AccountRepository) CommitTransaction(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}

func (r *AccountRepository) RollbackTransaction(ctx context.Context, tx pgx.Tx) error {
	return tx.Rollback(ctx)
}
