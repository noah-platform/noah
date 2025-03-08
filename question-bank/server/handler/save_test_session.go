package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/response"
	"github.com/noah-platform/noah/question-bank/server/core"
)

type SaveTestSessionRequest struct {
	Answers map[string]string `json:"answers" validate:"required"`
}

// SaveTestSession godoc
//
//	@Summary	Save test session
//	@Tags		external
//	@Router		/external/v1/tests/{testID} [put]
//	@Param		testID	path		string	true	"Test ID"
//	@Param		request	body	SaveTestSessionRequest	true	"Save test session request"
//	@Success	204 	"Test session saved"
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) SaveTestSession(c echo.Context) error {
	ctx := c.Request().Context()

	userID := s.auth.GetUserID(c)
	testID := c.Param("testID")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("userId", userID).Str("testId", testID).Logger()
	ctx = l.WithContext(ctx)

	var req SaveTestSessionRequest
	if err := c.Bind(&req); err != nil {
		l.Info().Err(err).Msg("[Server.SaveTestSession] failed to bind request")

		return response.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(req); err != nil {
		l.Info().Err(err).Msg("[Server.SaveTestSession] failed to validate request")

		return response.BadRequest(c, "invalid request body")
	}

	err := s.service.SaveTestSession(ctx, userID, testID, req.Answers)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrUserTestSessionNotFound):
			l.Info().Msg("[Server.SaveTestSession] test session not found")

			return response.NotFound(c, "test session not found")
		default:
			l.Error().Err(err).Msg("[Server.SaveTestSession] failed to save test session")

			return response.InternalServerError(c, "failed to save test session")
		}
	}

	l.Info().Msg("[Server.SaveTestSession] saved test session successfully")

	return response.Success(c)
}
