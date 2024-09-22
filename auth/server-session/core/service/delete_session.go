package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
)

func (s *Service) DeleteSession(ctx context.Context, id string) error {
	l := log.Ctx(ctx)

	err := s.sessionRepo.DeleteSession(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrSessionNotFound):
			l.Info().Err(err).Msg("[Service.DeleteSession] session not found")

			return core.ErrSessionNotFound
		default:
			l.Error().Err(err).Msg("[Service.DeleteSession] failed to delete session")

			return errors.Wrap(err, "failed to delete session")
		}
	}

	l.Info().Msg("[Service.DeleteSession] session deleted")

	return nil
}
