package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (s *Service) GetTestsByModule(ctx context.Context, module string) ([]core.TestInfo, error) {
	l := log.Ctx(ctx)

	testInfoList, err := s.questionBankRepo.GetTestsByModule(ctx, module)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTestNotFound):
			l.Info().Err(err).Msg("[Service.GetTestsByModule] example not found")

			return nil, core.ErrTestNotFound
		default:
			l.Error().Err(err).Msg("[Service.GetTestsByModule] failed to get tests by module")

			return nil, errors.Wrap(err, "failed to get example")
		}
	}

	l.Debug().Interface("tests", testInfoList).Msg("[Service.GetTestsByModule] got tests by module")

	return testInfoList, nil
}
