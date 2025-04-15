package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (s *Service) ListMyTestSessions(ctx context.Context, userID string) ([]core.UserTestSession, error) {
	l := log.Ctx(ctx)

	sessions, err := s.userTestSessionRepo.GetManySessionsByUserID(ctx, userID, true)
	if err != nil {
		l.Error().Err(err).Msg("[Service.ListMyTestSessions] failed to list my test sessions")

		return nil, errors.Wrap(err, "failed to list my test sessions")
	}

	return sessions, nil
}
