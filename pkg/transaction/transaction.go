package transaction

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Tx = pgx.Tx

type RepositoryWithTransaction interface {
	BeginTransaction(ctx context.Context) (pgx.Tx, error)
	CommitTransaction(ctx context.Context, tx pgx.Tx) error
	RollbackTransaction(ctx context.Context, tx pgx.Tx) error
}
