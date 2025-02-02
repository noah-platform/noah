package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/noah-platform/noah/account/server/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

func (r *AccountRepository) GetAccountByGoogleAccountID(ctx context.Context, googleAccountID string) (*core.Account, error) {
	l := log.Ctx(ctx)

	account, err := r.queries.GetAccountByGoogleAccountID(ctx, pgtype.Text{String: googleAccountID, Valid: true})
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			l.Info().Msg("[AccountRepository.GetAccountByGoogleAccountID] account not found")

			return nil, core.ErrAccountNotFound
		default:
			l.Error().Err(err).Msg("[AccountRepository.GetAccountByGoogleAccountID] failed to get account")

			return nil, errors.Wrap(err, "failed to get account")
		}
	}

	l.Info().Bool("isVerified", account.IsVerified).Time("updatedAt", account.UpdatedAt.Time).Msg("[AccountRepository.GetAccountByGoogleAccountID] account loaded")

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
