package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/response"
	"github.com/noah-platform/noah/question-bank/server/core"
)

type BeginTestSessionResponse struct {
	Test         *core.TestEntry   `bson:"test"`
	Answers      map[string]string `bson:"answers"`
	StartedAt    time.Time         `bson:"startedAt"`
	LastActiveAt time.Time         `bson:"lastActiveAt"`
	ElapsedTime  int               `bson:"elapsedTime"`
	CompletedAt  *time.Time        `bson:"completedAt"`
}

// BeginTestSession godoc
//
//	@Summary	Begin test session
//	@Tags		external
//	@Router		/external/v1/tests/{testID} [post]
//	@Param		testID	path		string	true	"Test ID"
//	@Success	200		{object}	response.DataResponse[BeginTestSessionResponse]
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) BeginTestSession(c echo.Context) error {
	ctx := c.Request().Context()

	userID := s.auth.GetUserID(c)
	testID := c.Param("testID")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("userId", userID).Str("testId", testID).Logger()
	ctx = l.WithContext(ctx)

	session, err := s.service.BeginTestSession(ctx, userID, testID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTestNotFound):
			l.Info().Msg("[Server.BeginTestSession] test not found")

			return response.NotFound(c, "test not found")
		default:
			l.Error().Err(err).Msg("[Server.BeginTestSession] failed to begin test session")

			return response.InternalServerError(c, "failed to begin test session")
		}
	}

	l.Info().Msg("[Server.BeginTestSession] begin test session successfully")

	return response.Ok(c, &BeginTestSessionResponse{
		Test:         session.Test,
		Answers:      session.Answers,
		StartedAt:    session.StartedAt,
		LastActiveAt: session.LastActiveAt,
		ElapsedTime:  session.ElapsedTime,
		CompletedAt:  session.CompletedAt,
	})
}
