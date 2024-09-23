package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/noah-platform/noah/pkg/response"
)

// Health godoc
//
//	@Summary	Check server status
//	@Tags		health
//	@Router		/health [get]
//	@Success	204
func (s *Server) Health(c echo.Context) error {
	return response.Success(c)
}
