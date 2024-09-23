package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/response"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type ConfirmPasswordResetRequest struct {
	Password string `json:"password" validate:"required,min=8,max=64"`
}

// ConfirmPasswordReset godoc
//
//	@Summary	Confirm password reset
//	@Tags		public
//	@Router		/v1/reset-password/{token} [post]
//	@Param		token	path	string	true	"Token"
//	@Param		request	body	ConfirmPasswordResetRequest	true	"Confirm Password Reset Request"
//	@Success	204		"Password reset successfully"
//	@Failure	400     {object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) ConfirmPasswordReset(c echo.Context) error {
	ctx := c.Request().Context()

	token := c.Param("token")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("token", token).Logger()
	ctx = l.WithContext(ctx)

	var req ConfirmPasswordResetRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.ConfirmPasswordReset] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.ConfirmPasswordReset] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	err := s.service.ConfirmPasswordReset(ctx, token, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTokenNotFound):
			l.Info().Msg("[Server.ConfirmPasswordReset] password reset token not found")

			return response.BadRequest(c, "invalid password reset token")
		case errors.Is(err, core.ErrTokenExpired):
			l.Info().Msg("[Server.ConfirmPasswordReset] password reset token expired")

			return response.Forbidden(c, "expired password reset token")
		default:
			l.Error().Err(err).Msg("[Server.ConfirmPasswordReset] failed to confirm password reset")

			return response.InternalServerError(c, "failed to confirm password reset")
		}
	}

	l.Info().Msg("[Server.ConfirmPasswordReset] password reset confirmed")

	return response.Success(c)
}
