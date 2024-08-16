package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github.com/noah-platform/noah/account/server/core"
)

func (r *AccountRepository) GetAccount(ctx context.Context, id string) (*core.Account, error) {
	l := log.Ctx(ctx)

	account, err := r.queries.GetAccount(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			l.Info().Msg("[AccountRepository.GetAccount] account not found")

			return nil, core.ErrAccountNotFound
		default:
			l.Error().Err(err).Msgf("[AccountRepository.GetAccount] failed to get account")

			return nil, errors.Wrap(err, "failed to get account")
		}
	}

	l.Info().Bool("isVerified", account.IsVerified).Time("updatedAt", account.UpdatedAt.Time).Msg("[AccountRepository.GetAccount] account loaded")

	return &core.Account{
		ID:              account.UserID,
		Email:           account.Email,
		Name:            account.Name,
		GoogleAccountID: lo.TernaryF(account.GoogleAccountID.Valid, func() *string { return &account.GoogleAccountID.String }, func() *string { return nil }),
		Password:        lo.TernaryF(account.Password.Valid, func() *string { return &account.Password.String }, func() *string { return nil }),
		IsVerified:      account.IsVerified,
		CreatedAt:       account.CreatedAt.Time,
		UpdatedAt:       account.UpdatedAt.Time,
	}, nil
}
