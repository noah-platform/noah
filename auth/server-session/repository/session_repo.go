package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"github.com/noah-platform/noah/auth/server-session/generated/sqlc"
)

type SessionRepository struct {
	db      *pgxpool.Pool
	queries *sqlc.Queries
	rdb     *redis.Client
}

type SessionRepoDependencies struct {
	PgxClient   *pgxpool.Pool
	RedisClient *redis.Client
}

func NewSessionRepository(deps SessionRepoDependencies) *SessionRepository {
	queries := sqlc.New(deps.PgxClient)

	return &SessionRepository{
		db:      deps.PgxClient,
		queries: queries,
		rdb:     deps.RedisClient,
	}
}
