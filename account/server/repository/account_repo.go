package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/noah-platform/noah/account/server/generated/sqlc"
)

type AccountRepository struct {
	queries *sqlc.Queries
}

type AccountRepoDependencies struct {
	PgxClient *pgxpool.Pool
}

func NewAccountRepository(deps AccountRepoDependencies) *AccountRepository {
	queries := sqlc.New(deps.PgxClient)

	return &AccountRepository{
		queries: queries,
	}
}
