package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/response"
)

type RegisterAccountRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

// RegisterAccount godoc
//
//	@Summary	Register new account
//	@Tags		public
//	@Router		/v1/register [post]
//	@Param		request	body	RegisterAccountRequest	true	"Register account request"
//	@Success	204		"Account registered, awaiting email verification"
//	@Failure	400     {object}	response.ErrorResponse
//	@Failure	409		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) RegisterAccount(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()
	ctx = l.WithContext(ctx)

	var req RegisterAccountRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.RegisterAccount] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.RegisterAccount] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	err := s.service.RegisterAccount(ctx, req.Email, req.Name, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountAlreadyExists):
			l.Info().Msg("[Server.RegisterAccount] account already exists")

			return response.Conflict(c, "account already exists")
		default:
			l.Error().Err(err).Msgf("[Server.RegisterAccount] failed to register account")

			return response.InternalServerError(c, "failed to register account")
		}
	}

	l.Info().Msg("[Server.RegisterAccount] account registered")

	return response.Success(c)
}
