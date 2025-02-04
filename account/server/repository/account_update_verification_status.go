package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/noah-platform/noah/account/server/generated/sqlc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *AccountRepository) UpdateAccountVerificationStatus(ctx context.Context, userID string, isVerified bool) error {
	l := log.Ctx(ctx)
	*l = l.With().Str("userId", userID).Bool("isVerified", isVerified).Logger()

	params := sqlc.UpdateAccountVerificationStatusParams{
		UserID:            userID,
		IsVerified:        isVerified,
		VerificationToken: pgtype.Text{Valid: false},
	}

	if err := r.queries.UpdateAccountVerificationStatus(ctx, params); err != nil {
		l.Error().Err(err).Msg("[AccountRepository.UpdateAccountVerificationStatus] failed to update account verification status")

		return errors.Wrap(err, "failed to update account verification status")
	}

	l.Info().Msg("[AccountRepository.UpdateAccountVerificationStatus] account verification status updated")

	return nil
}
