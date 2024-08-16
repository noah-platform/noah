package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/noah-platform/noah/pkg/response"
	"github.com/rs/zerolog/log"
)

type ResetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// ResetPassword godoc
//
//	@Summary	Request password reset
//	@Tags		public
//	@Router		/v1/reset-password [post]
//	@Param		request	body	ResetPasswordRequest	true	"Reset Password request"
//	@Success	204		"Password reset request processed"
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) ResetPassword(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()
	ctx = l.WithContext(ctx)

	var req ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.ResetPassword] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.ResetPassword] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	traceID := c.Response().Header().Get(echo.HeaderXRequestID)
	err := s.service.ResetPassword(ctx, traceID, req.Email)
	if err != nil {
		l.Error().Err(err).Msgf("[Server.ResetPassword] failed to reset password")

		return response.InternalServerError(c, "failed to reset password")
	}

	l.Info().Msg("[Server.ResetPassword] password reset request processed")

	return response.Success(c)
}
