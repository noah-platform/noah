package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/response"
)

type GetAccountResponse = core.Account

// InternalGetAccount godoc
//
//	@Summary	Get account by ID
//	@Tags		internal
//	@Router		/internal/v1/accounts/{userID} [get]
//	@Param		userID	path		string	true	"User ID"
//	@Success	200		{object}	response.DataResponse[GetAccountResponse]
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) InternalGetAccount(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Param("userID")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("userId", userID).Logger()
	ctx = l.WithContext(ctx)

	account, err := s.service.GetAccount(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			l.Info().Msg("[Server.InternalGetAccount] account not found")

			return response.NotFound(c, "account not found")
		default:
			l.Error().Err(err).Msgf("[Server.InternalGetAccount] failed to get account")

			return response.InternalServerError(c, "failed to get account")
		}
	}

	l.Info().Msg("[Server.InternalGetAccount] get account successfully")

	return response.Data(c, GetAccountResponse{
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
