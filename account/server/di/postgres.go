package di

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type PostgresConfig struct {
	DatabaseUrl string
}

func newPgxClient(cfg PostgresConfig) *pgxpool.Pool {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.DatabaseUrl)
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect to postgres")
	}
	// TODO: Close connection pool on shutdown
	// defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Panic().Err(err).Msg("failed to ping postgres")
	}

	return pool
}
