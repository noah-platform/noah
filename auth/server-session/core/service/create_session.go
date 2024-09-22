package service

import (
	"context"
	"net/netip"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
	"github.com/noah-platform/noah/pkg/random"
)

func (s *Service) CreateSession(ctx context.Context, userID string, ipAddress netip.Addr, userAgent string) (*core.Session, error) {
	l := log.Ctx(ctx)

	sessionID, err := random.GenerateRandomString(s.config.SessionIDLength)
	if err != nil {
		l.Error().Err(err).Msg("[Service.CreateSession] failed to generate sessionID")

		return nil, errors.Wrap(err, "failed to generate sessionID")
	}

	*l = l.With().Str("sessionId", sessionID).Str("userId", userID).Logger()
	ctx = l.WithContext(ctx)

	session, err := s.sessionRepo.CreateSession(ctx, &core.Session{
		SessionID: sessionID,
		UserID:    userID,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	})
	if err != nil {
		l.Error().Err(err).Msg("[Service.CreateSession] failed to create session")

		return nil, errors.Wrap(err, "failed to create session")
	}

	l.Info().Msg("[Service.CreateSession] session created")

	return session, nil
}
