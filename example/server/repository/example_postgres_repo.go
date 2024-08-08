package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/noah-platform/noah/example/server/generated/sqlc"
)

type ExamplePostgresRepository struct {
	queries *sqlc.Queries
}

type ExamplePostgresRepoDependencies struct {
	PgxClient *pgxpool.Pool
}

func NewExamplePostgresRepository(deps ExamplePostgresRepoDependencies) *ExamplePostgresRepository {
	queries := sqlc.New(deps.PgxClient)

	return &ExamplePostgresRepository{
		queries: queries,
	}
}
