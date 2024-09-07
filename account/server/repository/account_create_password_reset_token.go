package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/account/server/generated/sqlc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *AccountRepository) CreatePasswordResetToken(ctx context.Context, token core.PasswordResetToken) error {
	l := log.Ctx(ctx)
	*l = l.With().Str("token", token.Token).Str("userId", token.UserID).Logger()

	params := sqlc.CreatePasswordResetTokenParams{
		Token:     token.Token,
		UserID:    token.UserID,
		ExpiresAt: pgtype.Timestamp{Time: token.ExpiresAt, Valid: true},
	}

	if err := r.queries.CreatePasswordResetToken(ctx, params); err != nil {
		l.Error().Err(err).Msg("[AccountRepository.CreatePasswordResetToken] failed to create password reset token")

		return errors.Wrap(err, "failed to create password reset token")
	}

	l.Info().Msg("[AccountRepository.CreatePasswordResetToken] password reset token created")

	return nil
}
