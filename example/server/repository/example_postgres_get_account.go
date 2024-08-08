package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/example/server/core"
)

func (r *ExamplePostgresRepository) GetExample(ctx context.Context, id int) (*core.Example, error) {
	l := log.Ctx(ctx)

	example, err := r.queries.GetExample(ctx, int32(id))
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			l.Info().Msg("[ExamplePostgresRepository.GetExample] example not found")

			return nil, core.ErrExampleNotFound
		default:
			l.Error().Err(err).Msgf("[ExamplePostgresRepository.GetExample] failed to get example")

			return nil, errors.Wrap(err, "failed to get example")
		}
	}

	l.Debug().Msg("[ExamplePostgresRepository.GetExample] example loaded")

	return &core.Example{
		ID:    int(example.ExampleID),
		Title: example.Title,
	}, nil
}
