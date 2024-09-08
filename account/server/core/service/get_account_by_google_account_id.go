package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
)

func (s *Service) GetAccountByGoogleAccountID(ctx context.Context, googleAccountID string) (*core.Account, error) {
	l := log.Ctx(ctx)

	account, err := s.accountRepo.GetAccountByGoogleAccountID(ctx, googleAccountID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			l.Info().Err(err).Msg("[Service.GetAccountByGoogleAccountID] account not found")

			return nil, core.ErrAccountNotFound
		default:
			l.Error().Err(err).Msg("[Service.GetAccountByGoogleAccountID] failed to get account")

			return nil, errors.Wrap(err, "failed to get account")
		}
	}

	l.Debug().Interface("account", account).Msg("[Service.GetAccountByGoogleAccountID] got account")

	return account, nil
}
