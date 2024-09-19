package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
	"github.com/noah-platform/noah/pkg/response"
)

type GetSessionResponse = core.Session

// InternalGetSession godoc
//
//	@Summary	Get session by ID
//	@Tags		internal
//	@Router		/internal/v1/sessions/{sessionID} [get]
//	@Param		sessionID	path		string	true	"Session ID"
//	@Success	200		{object}	response.DataResponse[GetSessionResponse]
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) InternalGetSession(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("sessionID")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("sessionId", id).Logger()
	ctx = l.WithContext(ctx)

	session, err := s.service.GetSession(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrSessionNotFound):
			l.Info().Msg("[Server.InternalGetSession] session not found")

			return response.NotFound(c, "session not found")
		default:
			l.Error().Err(err).Msg("[Server.InternalGetSession] failed to get session")

			return response.InternalServerError(c, "failed to get session")
		}
	}

	l.Info().Msg("[Server.InternalGetSession] get session successfully")

	return response.Ok(c, GetSessionResponse{
		SessionID: session.SessionID,
		UserID:    session.UserID,
		IPAddress: session.IPAddress,
		UserAgent: session.UserAgent,
		CreatedAt: session.CreatedAt,
	})
}
