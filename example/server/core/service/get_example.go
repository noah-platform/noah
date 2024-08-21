package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/example/server/core"
)

func (s *Service) GetExample(ctx context.Context, id int) (*core.Example, error) {
	l := log.Ctx(ctx)

	example, err := s.exampleRepo.GetExample(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrExampleNotFound):
			l.Info().Err(err).Msg("[Service.GetExample] example not found")

			return nil, core.ErrExampleNotFound
		default:
			l.Error().Err(err).Msg("[Service.GetExample] failed to get example")

			return nil, errors.Wrap(err, "failed to get example")
		}
	}

	l.Debug().Interface("example", example).Msg("[Service.GetExample] got example")

	return example, nil
}
