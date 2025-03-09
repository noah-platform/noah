package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (s *Service) GetTestsByModule(ctx context.Context, userID, module string) ([]core.UserTestInfo, error) {
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

	sessions, err := s.userTestSessionRepo.GetManySessionsByUserID(ctx, userID)
	if err != nil {
		l.Error().Err(err).Msg("[Service.GetTestsByModule] failed to get user test sessions")

		return nil, errors.Wrap(err, "failed to get user test sessions")
	}
	sessionsMap := lo.Map(sessions, func(session core.UserTestSession, _ int) string {
		return session.Test.ID
	})
	userTestInfoList := lo.Map(testInfoList, func(test core.TestInfo, _ int) core.UserTestInfo {
		return core.UserTestInfo{
			ID:     test.ID,
			Module: test.Module,
			// TODO: Handle completed tests
			Status: lo.Ternary(lo.Contains(sessionsMap, test.ID), core.TestSessionStatusInProgress, core.TestSessionStatusNotStarted),
		}
	})

	l.Debug().Interface("tests", testInfoList).Msg("[Service.GetTestsByModule] got tests by module")

	return userTestInfoList, nil
}
