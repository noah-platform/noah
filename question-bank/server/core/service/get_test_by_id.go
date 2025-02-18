package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (s *Service) GetTestById(ctx context.Context, id string) (*core.TestEntry, error) {
	l := log.Ctx(ctx)

	test, err := s.questionBankRepo.GetTestById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTestNotFound):
			l.Info().Err(err).Msg("[Service.GetTestByID] test not found")

			return nil, core.ErrTestNotFound
		default:
			l.Error().Err(err).Msg("[Service.GetTestByID] failed to get test")

			return nil, errors.Wrap(err, "failed to get test")
		}
	}

	l.Debug().Interface("test", test).Msg("[Service.GetTestByID] got test")

	return test, nil
}
