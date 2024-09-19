package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
	"github.com/noah-platform/noah/pkg/response"
)

type VerifySessionRequest struct {
	SessionID string `json:"sessionId"`
}

type VerifySessionResponse struct {
	UserID string `json:"userId"`
}

// InternalVerifySession godoc
//
//	@Summary	Verify a session
//	@Tags		internal
//	@Router		/internal/v1/sessions/verify [post]
//	@Param		request	body	VerifySessionRequest	true	"Verify session request"
//	@Success	201		{object}	response.DataResponse[VerifySessionResponse]
//	@Failure	401		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) InternalVerifySession(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()
	ctx = l.WithContext(ctx)

	var req VerifySessionRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.InternalVerifySession] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	userID, err := s.service.VerifySession(ctx, req.SessionID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrSessionNotFound):
			l.Warn().Msg("[Server.InternalVerifySession] invalid session")

			return response.Unauthorized(c, "invalid session")
		default:
			l.Error().Err(err).Msg("[Server.InternalVerifySession] failed to verify session")

			return response.InternalServerError(c, "failed to verify session")
		}
	}

	l.Info().Str("userId", userID).Msg("[Server.InternalVerifySession] session verified")

	return response.Ok(c, VerifySessionResponse{
		UserID: userID,
	})
}
