package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
	"github.com/noah-platform/noah/pkg/response"
)

// InternalDeleteSession godoc
//
//	@Summary	Delete session by ID
//	@Tags		internal
//	@Router		/internal/v1/sessions/{sessionID} [delete]
//	@Param		sessionID	path		string	true	"Session ID"
//	@Success	204		"Session deleted"
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) InternalDeleteSession(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("sessionID")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("sessionId", id).Logger()
	ctx = l.WithContext(ctx)

	err := s.service.DeleteSession(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrSessionNotFound):
			l.Info().Msg("[Server.InternalDeleteSession] session not found")

			return response.NotFound(c, "session not found")
		default:
			l.Error().Err(err).Msg("[Server.InternalDeleteSession] failed to delete session")

			return response.InternalServerError(c, "failed to delete session")
		}
	}

	l.Info().Msg("[Server.InternalDeleteSession] delete session successfully")

	return response.Success(c)
}
