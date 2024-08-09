package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/account/server/generated/sqlc"
)

func (r *AccountRepository) CreateAccount(ctx context.Context, account core.Account) error {
	l := log.Ctx(ctx)
	*l = l.With().Str("userId", account.ID).
		Str("email", account.Email).
		Str("googleAccountId", lo.TernaryF(account.GoogleAccountID != nil, func() string { return *account.GoogleAccountID }, func() string { return "" })).
		Logger()

	params := sqlc.CreateAccountParams{
		UserID:          account.ID,
		Email:           account.Email,
		Name:            account.Name,
		Password:        pgtype.Text{Valid: false},
		GoogleAccountID: pgtype.Text{Valid: false},
		IsVerified:      account.IsVerified,
	}
	if account.GoogleAccountID != nil {
		params.GoogleAccountID = pgtype.Text{String: *account.GoogleAccountID, Valid: true}
	}
	if account.Password != nil {
		params.Password = pgtype.Text{String: *account.Password, Valid: true}
	}

	err := r.queries.CreateAccount(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			l.Info().Err(err).Msgf("[AccountRepository.CreateAccount] account already exists")

			return core.ErrAccountAlreadyExists
		default:
			l.Error().Err(err).Msgf("[AccountRepository.CreateAccount] failed to create account")

			return errors.Wrap(err, "failed to create account")
		}
	}

	l.Info().Msg("[AccountRepository.CreateAccount] account created")

	return nil
}
