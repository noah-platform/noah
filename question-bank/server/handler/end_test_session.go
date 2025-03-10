package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/response"
	"github.com/noah-platform/noah/question-bank/server/core"
)

type EndTestSessionRequest struct {
	Answers map[string]string `json:"answers" validate:"required"`
}

// EndTestSession godoc
//
//	@Summary	End test session
//	@Tags		external
//	@Router		/external/v1/tests/{testID}/end [post]
//	@Param		testID	path		string	true	"Test ID"
//	@Param		request	body	EndTestSessionRequest	true	"End test session request"
//	@Success	204 	"Test session ended"
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) EndTestSession(c echo.Context) error {
	ctx := c.Request().Context()

	userID := s.auth.GetUserID(c)
	testID := c.Param("testID")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("userId", userID).Str("testId", testID).Logger()
	ctx = l.WithContext(ctx)

	var req EndTestSessionRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.EndTestSession] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.EndTestSession] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	err := s.service.EndTestSession(ctx, userID, testID, req.Answers)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrUserTestSessionNotFound):
			l.Info().Msg("[Server.EndTestSession] test session not found")

			return response.NotFound(c, "test session not found")
		default:
			l.Error().Err(err).Msg("[Server.EndTestSession] failed to end test session")

			return response.InternalServerError(c, "failed to end test session")
		}
	}

	l.Info().Msg("[Server.EndTestSession] test session ended")

	return response.Success(c)
}
