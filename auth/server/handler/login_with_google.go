package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server/core"
	"github.com/noah-platform/noah/pkg/response"
)

type LoginWithGoogleRequest struct {
	IDToken string `json:"idToken" validate:"required"`
}

// LoginWithGoogle godoc
//
//	@Summary	Login with Google
//	@Tags		public
//	@Router		/v1/login/google [post]
//	@Param		request	body	LoginWithGoogleRequest	true	"Login with Google Account request"
//	@Success	204
//	@Failure	400		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) LoginWithGoogle(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()
	ctx = l.WithContext(ctx)

	var req LoginWithGoogleRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.LoginWithGoogle] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.LoginWithGoogle] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	sessionID, err := s.service.LoginWithGoogle(ctx, req.IDToken)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrInvalidCredentials):
			l.Info().Msg("[Server.LoginWithGoogle] invalid id token")

			return response.BadRequest(c, "invalid id token")
		default:
			l.Error().Err(err).Msg("[Server.LoginWithGoogle] failed to login")

			return response.InternalServerError(c, "failed to login")
		}
	}

	cookie := &http.Cookie{
		Name:     sessionCookieName,
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)

	l.Info().Str("sessionId", sessionID).Msg("[Server.LoginWithGoogle] login successfully")

	return response.Success(c)
}
