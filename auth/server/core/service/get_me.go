package service

import (
	"context"

	"github.com/noah-platform/noah/account/server/client"
	"github.com/noah-platform/noah/auth/server/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *Service) GetMe(ctx context.Context, userID string) (*core.Me, error) {
	l := log.Ctx(ctx)

	account, err := s.accountClient.FetchAccount(userID)
	if err != nil {
		switch {
		case errors.Is(err, client.ErrAccountNotFound):
			l.Error().Err(err).Msg("[Service.GetMe] account not found")

			return nil, core.ErrInvalidSession
		default:
			l.Error().Err(err).Msg("[Service.GetMe] failed to get account")

			return nil, errors.Wrap(err, "failed to get account")
		}
	}

	l.Info().Msg("[Service.GetMe] get me successfully")

	me := &core.Me{
		UserID:     account.ID,
		Name:       account.Name,
		Email:      account.Email,
		IsVerified: account.IsVerified,
	}
	return me, nil
}
