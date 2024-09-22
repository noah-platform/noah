package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
	"github.com/noah-platform/noah/pkg/response"
)

const sessionCookieName = "noahses"

const authRequestIdHeaderName = "X-Noah-Auth-Request-Id"
const userIdHeaderName = "X-Noah-User-Id"

func (s *Server) ExtAuth(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("path", c.Request().URL.Path).Logger()
	ctx = l.WithContext(ctx)

	sessionID, err := c.Cookie(sessionCookieName)
	if err != nil {
		l.Info().Err(err).Msg("[Server.ExtAuth] session cookie not found")

		return response.Unauthorized(c, "no session")
	}

	userID, err := s.service.VerifySession(ctx, sessionID.Value)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrSessionNotFound):
			l.Warn().Msg("[Server.ExtAuth] invalid session")

			// Expire the invalid session cookie
			cookie := &http.Cookie{
				Name:     sessionCookieName,
				Value:    "",
				Path:     "/",
				Expires:  time.Unix(0, 0),
				HttpOnly: true,
				Secure:   true,
			}
			c.SetCookie(cookie)

			return response.Unauthorized(c, "invalid session")
		default:
			l.Error().Err(err).Msg("[Server.ExtAuth] failed to verify session")

			return response.InternalServerError(c, "failed to verify session")
		}
	}

	c.Response().Header().Set(authRequestIdHeaderName, c.Response().Header().Get(echo.HeaderXRequestID))
	c.Response().Header().Set(userIdHeaderName, userID)

	l.Info().Str("userId", userID).Msg("[Server.ExtAuth] session verified")

	return response.Ok(c, "ok")
}
