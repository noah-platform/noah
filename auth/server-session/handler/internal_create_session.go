package handler

import (
	"net/netip"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/auth/server-session/core"
	"github.com/noah-platform/noah/pkg/response"
)

type CreateSessionRequest struct {
	UserID    string `json:"userId" validate:"required"`
	IPAddress string `json:"ipAddress" validate:"required,ip"`
	UserAgent string `json:"userAgent" validate:"required"`
}

type CreateSessionResponse = core.Session

// InternalCreateSession godoc
//
//	@Summary	Create new session
//	@Tags		internal
//	@Router		/internal/v1/sessions [post]
//	@Param		request	body	CreateSessionRequest	true	"Create session request"
//	@Success	201		{object}	response.DataResponse[CreateSessionResponse]
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) InternalCreateSession(c echo.Context) error {
	ctx := c.Request().Context()

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()
	ctx = l.WithContext(ctx)

	var req CreateSessionRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.InternalCreateSession] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.InternalCreateSession] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	ipAddress, err := netip.ParseAddr(req.IPAddress)
	if err != nil {
		l.Info().Err(err).Msg("[Server.InternalCreateSession] failed to parse IP address")

		return response.BadRequest(c, "invalid IP address")
	}

	session, err := s.service.CreateSession(ctx, req.UserID, ipAddress, req.UserAgent)
	if err != nil {
		l.Error().Err(err).Msg("[Server.InternalCreateSession] failed to create session")

		return response.InternalServerError(c, "failed to create session")
	}

	l.Info().Msg("[Server.InternalCreateSession] session created")

	return response.Created(c, CreateSessionResponse{
		SessionID: session.SessionID,
		UserID:    session.UserID,
		IPAddress: session.IPAddress,
		UserAgent: session.UserAgent,
		CreatedAt: session.CreatedAt,
	})
}
