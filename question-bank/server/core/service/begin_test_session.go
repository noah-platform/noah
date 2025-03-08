package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (s *Service) BeginTestSession(ctx context.Context, userID, testID string) (*core.UserTestSession, error) {
	l := log.Ctx(ctx)

	test, err := s.questionBankRepo.GetTestById(ctx, testID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTestNotFound):
			l.Info().Msg("[Service.BeginTestSession] test not found")

			return nil, core.ErrTestNotFound
		default:
			l.Error().Err(err).Msg("[Service.BeginTestSession] failed to get test from question bank")

			return nil, errors.Wrap(err, "failed to get test from question bank")
		}
	}

	session, err := s.userTestSessionRepo.GetSession(ctx, userID, testID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrUserTestSessionNotFound):
			break
		default:
			l.Error().Err(err).Msg("[Service.BeginTestSession] failed to get user test session")

			return nil, errors.Wrap(err, "failed to get user test session")
		}
	}

	if errors.Is(err, core.ErrUserTestSessionNotFound) {
		l.Info().Msg("[Service.BeginTestSession] user test session not found, creating new session")

		session = &core.UserTestSession{
			Test:    test,
			Answers: make(map[string]string),
		}
		err = s.userTestSessionRepo.SaveSession(ctx, userID, testID, session)
		if err != nil {
			l.Error().Err(err).Msg("[Service.BeginTestSession] failed to save user test session")

			return nil, errors.Wrap(err, "failed to save user test session")
		}
	} else {
		l.Info().Msg("[Service.BeginTestSession] existing user test session found, continuing previous session")
	}

	l.Info().Msg("[Service.BeginTestSession] begin test session successfully")

	return session, nil
}
