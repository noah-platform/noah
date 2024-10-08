package service

import (
	"context"
	"time"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) ConfirmPasswordReset(ctx context.Context, token, password string) (err error) {
	l := log.Ctx(ctx)
	*l = l.With().Str("token", token).Logger()
	ctx = l.WithContext(ctx)

	// get account using token > just GetPasswordResetToken
	passwordResetToken, err := s.accountRepo.GetPasswordResetToken(ctx, token)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTokenNotFound):
			l.Info().Msg("[Service.ConfirmPasswordReset] password reset token not found")

			return core.ErrTokenNotFound
		default:
			l.Error().Err(err).Msg("[Service.ConfirmPasswordReset] failed to get password reset token")

			return errors.Wrap(err, "failed to get password reset token")
		}
	}

	// validate token, check if token not expired
	if passwordResetToken.ExpiresAt.Before(time.Now()) {
		l.Info().Msg("[Service.ConfirmPasswordReset] password reset token expired")

		return core.ErrTokenExpired
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		l.Error().Err(err).Msg("[Service.ConfirmPasswordReset] failed to hash password")

		return errors.Wrap(err, "failed to hash password")
	}
	passwordHash := string(hash)

	// add transaction: not sure if this is needed, and how to implement it
	// update account password
	err = s.accountRepo.UpdateAccountPassword(ctx, passwordResetToken.UserID, passwordHash)
	if err != nil {
		l.Error().Err(err).Msg("[Service.ConfirmPasswordReset] failed to update account password")

		return errors.Wrap(err, "failed to update account password")
	}

	// invalidate token > just DeletePasswordResetToken and log
	err = s.accountRepo.DeletePasswordResetToken(ctx, token)
	if err != nil {
		l.Error().Err(err).Msg("[Service.ConfirmPasswordReset] failed to delete password reset token")

		return errors.Wrap(err, "failed to delete password reset token")
	}

	l.Info().Msg("[Service.ConfirmPasswordReset] password reset confirmed")

	return nil
}
