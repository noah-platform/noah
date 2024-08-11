package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
)

func (s *Service) GetSession(ctx context.Context, id string) (*core.Session, error) {
	l := log.Ctx(ctx)

	session, err := s.sessionRepo.GetSession(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrSessionNotFound):
			l.Info().Err(err).Msg("[Service.GetSession] session not found")

			return nil, core.ErrSessionNotFound
		default:
			l.Error().Err(err).Msg("[Service.GetSession] failed to get session")

			return nil, errors.Wrap(err, "failed to get session")
		}
	}

	l.Debug().Interface("session", session).Msg("[Service.GetSession] got session")

	return session, nil
}
