package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *AccountRepository) DeletePasswordResetToken(ctx context.Context, token string) error {
	l := log.Ctx(ctx)
	*l = l.With().Str("token", token).Logger()

	if err := r.queries.DeletePasswordResetToken(ctx, token); err != nil {
		l.Error().Err(err).Msg("[AccountRepository.DeletePasswordResetToken] failed to delete password reset token")

		return errors.Wrap(err, "failed to delete password reset token")
	}

	l.Info().Msg("[AccountRepository.DeletePasswordResetToken] password reset token deleted")

	return nil
}
