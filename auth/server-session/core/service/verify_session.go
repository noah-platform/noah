package service

import (
	"context"

	"github.com/noah-platform/noah/auth/server-session/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *Service) VerifySession(ctx context.Context, id string) (string, error) {
	l := log.Ctx(ctx)

	if id == "" || len(id) != s.config.SessionIDLength {
		l.Warn().Msg("[Service.VerifySession] invalid session format")

		return "", core.ErrSessionNotFound
	}

	userID, err := s.sessionRepo.GetUserIDFromSession(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrSessionNotFound):
			l.Info().Err(err).Msg("[Service.VerifySession] session not found")

			return "", core.ErrSessionNotFound
		default:
			l.Error().Err(err).Msg("[Service.VerifySession] failed to get session")

			return "", errors.Wrap(err, "failed to get session")
		}
	}

	l.Info().Str("userId", userID).Msg("[Service.VerifySession] session verified")

	return userID, nil
}
