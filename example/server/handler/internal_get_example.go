package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/example/server/core"
	"github.com/noah-platform/noah/pkg/response"
)

type GetExampleResponse = core.Example

// InternalGetExample godoc
//
//	@Summary	Get example by ID
//	@Tags		internal
//	@Router		/internal/v1/example/{exampleID} [get]
//	@Param		exampleID	path		int	true	"Example ID"
//	@Success	200		{object}	response.DataResponse[GetExampleResponse]
//	@Failure	404		{object}	response.ErrorResponse
//	@Failure	500		{object}	response.ErrorResponse
func (s *Server) InternalGetExample(c echo.Context) error {
	ctx := c.Request().Context()

	exampleID := c.Param("exampleID")

	id, err := strconv.Atoi(exampleID)
	if err != nil {
		log.Info().Err(err).Msg("[Server.InternalGetExample] failed to convert exampleID to int")

		return response.BadRequest(c, "exampleID must be an integer")
	}

	l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Int("exampleID", id).Logger()
	ctx = l.WithContext(ctx)

	example, err := s.service.GetExample(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrExampleNotFound):
			l.Info().Msg("[Server.InternalGetExample] example not found")

			return response.NotFound(c, "example not found")
		default:
			l.Error().Err(err).Msg("[Server.InternalGetExample] failed to get example")

			return response.InternalServerError(c, "failed to get example")
		}
	}

	l.Info().Msg("[Server.InternalGetExample] get example successfully")

	return response.Data(c, GetExampleResponse{
		ID:    example.ID,
		Title: example.Title,
	})
}
