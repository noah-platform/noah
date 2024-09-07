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

	// call auth session server to delete session

	l.Info().Msg("[Service.Logout] logout successfully")

	return nil
}
