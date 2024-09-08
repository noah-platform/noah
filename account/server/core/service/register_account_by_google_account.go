package service

import (
	"context"
	"crypto/rand"

	"github.com/lucsky/cuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
)

func (s *Service) RegisterAccountByGoogleAccount(ctx context.Context, email, name, googleAccountID string) (*core.Account, error) {
	l := log.Ctx(ctx)
	*l = l.With().Str("email", email).Str("googleAccountId", googleAccountID).Logger()
	ctx = l.WithContext(ctx)

	userID, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		l.Error().Err(err).Msg("[Service.RegisterAccountByGoogleAccount] failed to generate userId")

		return nil, errors.Wrap(err, "failed to generate userId")
	}

	err = s.accountRepo.CreateAccount(ctx, nil, core.Account{
		ID:              userID,
		Email:           email,
		Name:            name,
		GoogleAccountID: &googleAccountID,
		IsVerified:      true,
	})
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountAlreadyExists):
			l.Info().Msg("[Service.RegisterAccountByGoogleAccount] account already exists")

			return nil, errors.Wrap(err, "account already exists")

		default:
			l.Error().Err(err).Msg("[Service.RegisterAccountByGoogleAccount] failed to create account")

			return nil, errors.Wrap(err, "failed to create account")
		}
	}

	account, err := s.GetAccountByGoogleAccountID(ctx, googleAccountID)
	if err != nil {
		l.Error().Err(err).Msg("[Service.RegisterAccountByGoogleAccount] failed to get account by google account id")

		return nil, errors.Wrap(err, "failed to get account by google account id")
	}

	l.Info().Msg("[Service.RegisterAccountByGoogleAccount] account created")

	return account, nil
}
