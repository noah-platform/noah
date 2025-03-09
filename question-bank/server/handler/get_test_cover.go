package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/response"
	"github.com/noah-platform/noah/question-bank/server/core"
)

type GetTestCoverResponse = core.TestCover

// GetTestCover godoc
//
//	@Summary	Get test cover
//	@Tags		external
//	@Router		/external/v1/tests/{testID} [get]
//	@Param		testID	path		string	true	"Test ID"
//	@Success	200		{object}	response.DataResponse[GetTestCoverResponse]
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) GetTestCover(c echo.Context) error {
	ctx := c.Request().Context()

	testID := c.Param("testID")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("testID", testID).Logger()
	ctx = l.WithContext(ctx)

	test, err := s.service.GetTestCover(ctx, testID)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTestNotFound):
			l.Info().Msg("[Server.GetTestCover] test not found")

			return response.NotFound(c, "test not found")
		default:
			l.Error().Err(err).Msg("[Server.GetTestCover] failed to get test cover")

			return response.InternalServerError(c, "failed to get test cover")
		}
	}

	l.Info().Msg("[Server.GetTestCover] get test successfully")

	return response.Ok(c, GetTestCoverResponse{
		ID:          test.ID,
		Title:       test.Title,
		Duration:    test.Duration,
		Instruction: test.Instruction,
		Module:      test.Module,
		Tags:        test.Tags,
	})
}
