package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server/core"
	"github.com/noah-platform/noah/pkg/response"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login godoc
//
//	@Summary	Login
//	@Tags		public
//	@Router		/v1/login [post]
//	@Param		request	body	LoginRequest	true	"Login request"
//	@Success	204
//	@Failure	400		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) Login(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()
	ctx = l.WithContext(ctx)

	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.Login] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.Login] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	l = l.With().Str("email", req.Email).Logger()

	sessionID, err := s.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrInvalidCredentials):
			l.Info().Msg("[Server.Login] invalid email or password")

			return response.BadRequest(c, "invalid email or password")
		case errors.Is(err, core.ErrAccountNotVerified):
			l.Info().Msg("[Server.Login] account not verified")

			return response.BadRequest(c, "account not verified")
		default:
			l.Error().Err(err).Msg("[Server.Login] failed to login")

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

	l.Info().Str("sessionId", sessionID).Msg("[Server.Login] login successfully")

	return response.Success(c)
}
