package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
)

func (s *Service) GetAccount(ctx context.Context, id string) (*core.Account, error) {
	l := log.Ctx(ctx)

	account, err := s.accountRepo.GetAccount(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			l.Info().Err(err).Msg("[Service.GetAccount] account not found")

			return nil, core.ErrAccountNotFound
		default:
			l.Error().Err(err).Msgf("[Service.GetAccount] failed to get account")

			return nil, errors.Wrap(err, "failed to get account")
		}
	}

	l.Debug().Interface("account", account).Msg("[Service.GetAccount] got account")

	return account, nil
}
