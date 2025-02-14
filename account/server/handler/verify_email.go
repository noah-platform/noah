package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/response"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type VerifyEmailRequest struct {
	Token string `json:"token" validate:"required"`
}

// VerifyEmail godoc
//
//	@Summary	Verify email
//	@Tags		public
//	@Router		/external/v1/verify-email [post]
//	@Param		request	body	VerifyEmailRequest	true	"Verify Email Request"
//	@Success	204		"Email verified"
//	@Failure	400     {object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) VerifyEmail(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()
	ctx = l.WithContext(ctx)

	var req VerifyEmailRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.VerifyEmail] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.VerifyEmail] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	err := s.service.VerifyEmail(ctx, req.Token)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTokenNotFound):
			l.Info().Msg("[Server.VerifyEmail] invalid token")

			return response.BadRequest(c, "invalid token")
		default:
			l.Error().Err(err).Msg("[Server.VerifyEmail] failed to verify email")

			return response.InternalServerError(c, "failed to verify email")
		}
	}

	l.Info().Msg("[Server.VerifyEmail] email verified")

	return response.Success(c)
}
