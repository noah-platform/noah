package handler

import (
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/response"
)

// Docs godoc
//
//	@title			auth-server-session
//	@version		1.0.0
//	@host			localhost:8080
//	@contact.name	Noah Platform
//	@contact.email	noah-platform@googlegroups.com
//	@license.name	Proprietary
func (s *Server) Docs(c echo.Context) error {
	content, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: "./generated/docs/swagger.json",
	})

	if err != nil {
		log.Error().Err(err).Msg("[Server.Docs] failed to serve docs")

		return response.Error(c, http.StatusInternalServerError, "failed to serve docs")
	}

	return c.HTML(http.StatusOK, content)
}
