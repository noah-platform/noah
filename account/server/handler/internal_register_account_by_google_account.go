package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/response"
)

type RegisterAccountByGoogleAccountRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Name            string `json:"name" validate:"required"`
	GoogleAccountID string `json:"googleAccountId" validate:"required"`
}

type RegisterAccountByGoogleAccountResponse = core.Account

// InternalRegisterAccountByGoogleAccount godoc
//
//	@Summary	Register account by Google Account
//	@Tags		internal
//	@Router		/internal/v1/accounts/google [post]
//	@Param		request	body	RegisterAccountByGoogleAccountRequest	true	"Register account by Google Account request"
//	@Success	200		{object}	response.DataResponse[RegisterAccountByGoogleAccountResponse]
//	@Failure	400     {object}	response.ErrorResponse
//	@Failure	409		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) InternalRegisterAccountByGoogleAccount(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()
	ctx = l.WithContext(ctx)

	var req RegisterAccountByGoogleAccountRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.InternalRegisterAccountByGoogleAccount] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.InternalRegisterAccountByGoogleAccount] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	account, err := s.service.RegisterAccountByGoogleAccount(ctx, req.Email, req.Name, req.GoogleAccountID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountAlreadyExists):
			l.Info().Msg("[Server.InternalRegisterAccountByGoogleAccount] account already exists")

			return response.Conflict(c, "account already exists")
		default:
			l.Error().Err(err).Msg("[Server.InternalRegisterAccountByGoogleAccount] failed to register account")

			return response.InternalServerError(c, "failed to register account")
		}
	}

	l.Info().Msg("[Server.InternalRegisterAccountByGoogleAccount] account registered")

	return response.Data(c, RegisterAccountByGoogleAccountResponse{
		ID:              account.ID,
		Email:           account.Email,
		Name:            account.Name,
		GoogleAccountID: account.GoogleAccountID,
		Password:        account.Password,
		IsVerified:      account.IsVerified,
		CreatedAt:       account.CreatedAt,
		UpdatedAt:       account.UpdatedAt,
	})
}
