package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/response"
)

type GetAccountByGoogleAccountIDResponse = core.Account

// InternalGetAccountByGoogleAccountID godoc
//
//	@Summary	Get account by Google Account ID
//	@Tags		internal
//	@Router		/internal/v1/accounts/google/{googleAccountID} [get]
//	@Param		googleAccountID	path		string	true	"Google Account ID"
//	@Success	200		{object}	response.DataResponse[GetAccountByGoogleAccountIDResponse]
//	@Failure	400		{object}	response.ErrorResponse
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) InternalGetAccountByGoogleAccountID(c echo.Context) error {
	ctx := c.Request().Context()

	googleAccountID := c.Param("googleAccountID")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("googleAccountId", googleAccountID).Logger()
	ctx = l.WithContext(ctx)

	account, err := s.service.GetAccountByGoogleAccountID(ctx, googleAccountID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			l.Info().Msg("[Server.InternalGetAccountByGoogleAccountID] account not found")

			return response.NotFound(c, "account not found")
		default:
			l.Error().Err(err).Msg("[Server.InternalGetAccountByGoogleAccountID] failed to get account")

			return response.InternalServerError(c, "failed to get account")
		}
	}

	l.Info().Msg("[Server.InternalGetAccountByGoogleAccountID] get account successfully")

	return response.Data(c, GetAccountByGoogleAccountIDResponse{
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
