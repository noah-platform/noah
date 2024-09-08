package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/response"
)

type GetAccountByEmailResponse = core.Account

// InternalGetAccountByEmail godoc
//
//	@Summary	Get account by email
//	@Tags		internal
//	@Router		/internal/v1/accounts/email/{email} [get]
//	@Param		email	path		string	true	"Email"
//	@Success	200		{object}	response.DataResponse[GetAccountByEmailResponse]
//	@Failure	400		{object}	response.ErrorResponse
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) InternalGetAccountByEmail(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.Param("email")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("email", email).Logger()
	ctx = l.WithContext(ctx)

	account, err := s.service.GetAccountByEmail(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			l.Info().Msg("[Server.InternalGetAccountByEmail] account not found")

			return response.NotFound(c, "account not found")
		default:
			l.Error().Err(err).Msg("[Server.InternalGetAccountByEmail] failed to get account")

			return response.InternalServerError(c, "failed to get account")
		}
	}

	l.Info().Msg("[Server.InternalGetAccountByEmail] get account successfully")

	return response.Data(c, GetAccountByEmailResponse{
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
