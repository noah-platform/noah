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

func (r *AccountRepository) GetAccountByVerificationToken(ctx context.Context, token string) (*core.Account, error) {
	l := log.Ctx(ctx)

	account, err := r.queries.GetAccountByVerificationToken(ctx, pgtype.Text{String: token, Valid: true})
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			l.Info().Msg("[AccountRepository.GetAccountByVerificationToken] account not found")

			return nil, core.ErrAccountNotFound
		default:
			l.Error().Err(err).Msg("[AccountRepository.GetAccountByVerificationToken] failed to get account")

			return nil, errors.Wrap(err, "failed to get account")
		}
	}

	l.Info().Msg("[AccountRepository.GetAccountByVerificationToken] account loaded")

	return &core.Account{
		ID:                account.UserID,
		Email:             account.Email,
		Name:              account.Name,
		GoogleAccountID:   lo.TernaryF(account.GoogleAccountID.Valid, func() *string { return &account.GoogleAccountID.String }, func() *string { return nil }),
		Password:          lo.TernaryF(account.Password.Valid, func() *string { return &account.Password.String }, func() *string { return nil }),
		IsVerified:        account.IsVerified,
		VerificationToken: account.VerificationToken.String,
		CreatedAt:         account.CreatedAt.Time,
		UpdatedAt:         account.UpdatedAt.Time,
	}, nil
}
