package service

import (
	"context"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *Service) VerifyEmail(ctx context.Context, token string) (err error) {
	l := log.Ctx(ctx)

	*l = l.With().Str("token", token).Logger()
	ctx = l.WithContext(ctx)

	account, err := s.accountRepo.GetAccountByVerificationToken(ctx, token)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			l.Info().Err(err).Msg("[Service.VerifyEmail] invalid token")

			return core.ErrTokenNotFound
		default:
			l.Error().Err(err).Msgf("[Service.VerifyEmail] failed to get account by verification token")

			return errors.Wrap(err, "failed to get account by verification token")
		}
	}

	if err := s.accountRepo.UpdateAccountVerificationStatus(ctx, account.ID, true); err != nil {
		l.Error().Err(err).Msg("[Service.VerifyEmail] failed to update account verification status")

		return errors.Wrap(err, "failed to update account verification status")
	}

	l.Info().Str("userID", account.ID).Str("email", account.Email).Msg("[Service.VerifyEmail] email verified")

	return nil
}
