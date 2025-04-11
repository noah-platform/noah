package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github.com/noah-platform/noah/pkg/response"
	"github.com/noah-platform/noah/question-bank/server/core"
)

type MyTestSession struct {
	Test         *core.TestEntry   `json:"test"`
	Answers      map[string]string `json:"answers"`
	StartedAt    time.Time         `json:"startedAt"`
	LastActiveAt time.Time         `json:"lastActiveAt"`
	ElapsedTime  int               `json:"elapsedTime"`
	CompletedAt  *time.Time        `json:"completedAt"`
}

type ListMyTestSessionsResponse = []MyTestSession

// ListMyTestSessionsResponse godoc
//
//	@Summary	List my test sessions
//	@Tags		external
//	@Router		/external/v1/sessions [get]
//	@Success	200		{object}	response.DataResponse[ListMyTestSessionsResponse]
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) ListMyTestSessions(c echo.Context) error {
	ctx := c.Request().Context()

	userID := s.auth.GetUserID(c)

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("userID", userID).Logger()
	ctx = l.WithContext(ctx)

	sessions, err := s.service.ListMyTestSessions(ctx, userID)
	if err != nil {
		l.Error().Err(err).Msg("[Server.ListMyTestSessions] failed to list my test sessions")

		return response.InternalServerError(c, "failed to list my test sessions")
	}

	l.Info().Msg("[Server.ListMyTestSessions] list my tests sessions successfully")

	mySessions := lo.Map(sessions, func(session core.UserTestSession, _ int) MyTestSession {
		return MyTestSession{
			Test:         session.Test,
			Answers:      session.Answers,
			StartedAt:    session.StartedAt,
			LastActiveAt: session.LastActiveAt,
			ElapsedTime:  session.ElapsedTime,
			CompletedAt:  session.CompletedAt,
		}
	})
	return response.Ok(c, ListMyTestSessionsResponse(mySessions))
}
