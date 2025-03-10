package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (s *Service) GetTestCover(ctx context.Context, id string) (*core.TestCover, error) {
	l := log.Ctx(ctx)

	test, err := s.questionBankRepo.GetTestById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTestNotFound):
			l.Info().Err(err).Msg("[Service.GetTestCover] test not found")

			return nil, core.ErrTestNotFound
		default:
			l.Error().Err(err).Msg("[Service.GetTestCover] failed to get test")

			return nil, errors.Wrap(err, "failed to get test")
		}
	}

	l.Debug().Interface("test", test).Msg("[Service.GetTestCover] got test")

	return &core.TestCover{
		ID:          test.ID,
		Title:       test.Title,
		Duration:    test.Duration,
		Instruction: test.Instruction,
		Module:      test.Module,
		Tags:        test.Tags,
	}, nil
}
