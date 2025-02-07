package service

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (s *Service) Logout(ctx context.Context, sessionID string) error {
	l := log.Ctx(ctx)
	*l = l.With().Str("sessionId", sessionID).Logger()

	if sessionID == "" {
		l.Info().Msg("[Service.Logout] session ID is empty, assumed logged out")

		return nil
	}

	if err := s.authSessionClient.DeleteSession(sessionID); err != nil {
		l.Error().Err(err).Msg("[Service.Logout] failed to delete session")

		return err
	}

	l.Info().Msg("[Service.Logout] logout successfully")

	return nil
}
