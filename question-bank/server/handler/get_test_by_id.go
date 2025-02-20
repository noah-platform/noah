package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/response"
	"github.com/noah-platform/noah/question-bank/server/core"
)

type GetTestByIdResponse = core.TestEntry

// GetTestById godoc
//
//	@Summary	Get test by ID
//	@Tags		external
//	@Router		/external/v1/tests/{testID} [get]
//	@Param		testID	path		string	true	"Test ID"
//	@Success	200		{object}	response.DataResponse[GetTestByIdResponse]
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) GetTestById(c echo.Context) error {
	ctx := c.Request().Context()

	testID := c.Param("testID")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("testID", testID).Logger()
	ctx = l.WithContext(ctx)

	test, err := s.service.GetTestById(ctx, testID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTestNotFound):
			l.Info().Msg("[Server.GetTestById] test not found")

			return response.NotFound(c, "test not found")
		default:
			l.Error().Err(err).Msg("[Server.GetTestById] failed to get test")

			return response.InternalServerError(c, "failed to get test")
		}
	}

	l.Info().Msg("[Server.GetTestById] get test successfully")

	return response.Ok(c, GetTestByIdResponse{
		ID:          test.ID,
		Module:      test.Module,
		QuestionSet: test.QuestionSet,
		Audio:       test.Audio,
		AudioScript: test.AudioScript,
	})
}
