package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/noah-platform/noah/account/server/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *AccountRepository) GetPasswordResetToken(ctx context.Context, token string) (*core.PasswordResetToken, error) {
	l := log.Ctx(ctx)

	// get password reset token
	passwordResetToken, err := r.queries.GetPasswordResetToken(ctx, token)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			l.Info().Msg("[AccountRepository.GetPasswordResetToken] password reset token not exist")

			return nil, core.ErrTokenNotFound
		default:
			l.Error().Err(err).Msg("[AccountRepository.GetPasswordResetToken] failed to get password reset token")

			return nil, errors.Wrap(err, "failed to get password reset token")
		}
	}

	l.Info().Time("expiresAt", passwordResetToken.ExpiresAt.Time).Msg("[AccountRepository.GetPasswordResetToken] password reset token loaded")

	return &core.PasswordResetToken{
		Token:     passwordResetToken.Token,
		UserID:    passwordResetToken.UserID,
		CreatedAt: passwordResetToken.CreatedAt.Time,
		ExpiresAt: passwordResetToken.ExpiresAt.Time,
	}, nil
}
