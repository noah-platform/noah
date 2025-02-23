package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/response"
)

type GetMeResponse struct {
	UserID     string `json:"userId"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	IsVerified bool   `json:"isVerified"`
}

// Me godoc
//
//	@Summary	Get me
//	@Tags		external
//	@Router		/external/v1/me [get]
//	@Success	200		{object}	response.DataResponse[GetMeResponse]
//	@Failure	401		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) GetMe(c echo.Context) error {
	ctx := c.Request().Context()

	userID := s.auth.GetUserID(c)

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("userID", userID).Logger()
	ctx = l.WithContext(ctx)

	me, err := s.service.GetMe(ctx, userID)
	if err != nil {
		l.Error().Err(err).Msg("[Server.GetMe] failed to get me")

		return response.InternalServerError(c, "failed to get me")
	}

	l.Info().Msg("[Server.GetMe] get me successfully")

	return response.Ok(c, GetMeResponse{
		UserID:     me.UserID,
		Email:      me.Email,
		Name:       me.Name,
		IsVerified: me.IsVerified,
	})
}
