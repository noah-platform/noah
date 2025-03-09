package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (s *Service) SaveTestSession(ctx context.Context, userID, testID string, answers map[string]string) error {
	l := log.Ctx(ctx)

	session, err := s.userTestSessionRepo.GetSession(ctx, userID, testID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrUserTestSessionNotFound):
			l.Error().Msg("[Service.SaveTestSession] test session not found")

			return core.ErrUserTestSessionNotFound
		default:
			l.Error().Err(err).Msg("[Service.SaveTestSession] failed to get test session")

			return errors.Wrap(err, "failed to get test session")
		}
	}

	now := time.Now()
	session.Answers = answers
	session.ElapsedTime = session.ElapsedTime + int(now.Sub(session.LastActiveAt).Seconds())
	session.LastActiveAt = now
	err = s.userTestSessionRepo.SaveSession(ctx, session)
	if err != nil {
		l.Error().Err(err).Msg("[Service.SaveTestSession] failed to save test session")

		return errors.Wrap(err, "failed to save test session")
	}

	l.Info().Msg("[Service.SaveTestSession] saved test session successfully")

	return nil
}
