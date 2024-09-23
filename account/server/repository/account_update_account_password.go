package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/noah-platform/noah/account/server/generated/sqlc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *AccountRepository) UpdateAccountPassword(ctx context.Context, userId, password string) error {
	l := log.Ctx(ctx)
	*l = l.With().Str("userId", userId).Logger()

	params := sqlc.UpdateAccountPasswordParams{
		UserID:   userId,
		Password: pgtype.Text{String: password, Valid: true},
	}

	if err := r.queries.UpdateAccountPassword(ctx, params); err != nil {
		l.Error().Err(err).Msg("[AccountRepository.UpdateAccountPassword] failed to update account password")

		return errors.Wrap(err, "failed to update account password")
	}

	l.Info().Msg("[AccountRepository.UpdateAccountPassword] account password updated")

	return nil
}
