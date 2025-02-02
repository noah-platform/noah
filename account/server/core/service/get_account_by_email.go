package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
)

func (s *Service) GetAccountByEmail(ctx context.Context, email string) (*core.Account, error) {
	l := log.Ctx(ctx)

	account, err := s.accountRepo.GetAccountByEmail(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			l.Info().Err(err).Msg("[Service.GetAccountByEmail] account not found")

			return nil, core.ErrAccountNotFound
		default:
			l.Error().Err(err).Msgf("[Service.GetAccountByEmail] failed to get account")

			return nil, errors.Wrap(err, "failed to get account")
		}
	}

	l.Debug().Interface("account", account).Msg("[Service.GetAccountByEmail] got account")

	return account, nil
}
