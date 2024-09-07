package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *AccountRepository) ValidatePasswordResetToken(ctx context.Context, token string) error {
	l := log.Ctx(ctx)

	// check if token exists
	passwordResetToken, err := r.queries.GetPasswordResetToken(ctx, token)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			l.Info().Msg("[AccountRepository.ValidatePasswordResetToken] password reset token not exist")

			return errors.New("password reset token does not exist")
		default:
			l.Error().Err(err).Msg("[AccountRepository.ValidatePasswordResetToken] failed to get password reset token")

			return errors.Wrap(err, "failed to get password reset token")
		}
	}

	// check if existing token not expired
	if passwordResetToken.ExpiresAt.Time.Before(time.Now()) {
		l.Info().Msg("[AccountRepository.ValidatePasswordResetToken] password reset token expired")

		return errors.New("password reset token expired")
	}

	l.Info().Msg("[AccountRepository.ValidatePasswordResetToken] password reset token valid")

	return nil
}
