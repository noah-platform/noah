package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/response"
)

// Logout godoc
//
//	@Summary	Logout
//	@Tags		public
//	@Router		/v1/logout [post]
//	@Success	204
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) Logout(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()
	ctx = l.WithContext(ctx)

	sessionID, err := c.Cookie(sessionCookieName)
	if err != nil {
		l.Warn().Err(err).Msg("[Server.Logout] session cookie not found, assumed logged out")

		return response.Success(c)
	}

	l = l.With().Str("sessionId", sessionID.Value).Logger()

	err = s.service.Logout(ctx, sessionID.Value)
	if err != nil {
		l.Error().Err(err).Msg("[Server.Logout] failed to logout, continuing to clear cookie")
	}

	cookie := &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)

	l.Info().Msg("[Server.Logout] logout successfully")

	return response.Success(c)
}
