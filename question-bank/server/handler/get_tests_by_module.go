package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/response"
	"github.com/noah-platform/noah/question-bank/server/core"
)

type GetTestsByModuleResponse = []core.TestInfo

// GetTestsByModule godoc
//
//	@Summary	Get tests info by module
//	@Tags		external
//	@Router		/external/v1/tests [get]
//	@Param		module	query		string	true	"Module"
//	@Success	200		{object}	response.DataResponse[GetTestsByModuleResponse]
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) GetTestsByModule(c echo.Context) error {
	ctx := c.Request().Context()

	module := c.QueryParam("module")

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Str("module", module).Logger()
	ctx = l.WithContext(ctx)

	testInfoList, err := s.service.GetTestsByModule(ctx, module)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrTestNotFound):
			l.Info().Msg("[Server.GetTestsByModule] tests not found")

			return response.NotFound(c, "tests not found")
		default:
			l.Error().Err(err).Msg("[Server.GetTestsByModule] failed to get tests")

			return response.InternalServerError(c, "failed to get tests")
		}
	}

	l.Info().Msg("[Server.GetTestsByModule] get tests successfully")

	return response.Ok(c, GetTestsByModuleResponse(testInfoList))
}
