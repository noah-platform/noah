package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DataResponse[T any] struct {
	Data T `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Data[T any](c echo.Context, data T) error {
	res := DataResponse[T]{
		Data: data,
	}
	return c.JSON(http.StatusOK, res)
}

func Success(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func Error(c echo.Context, status int, msg string) error {
	res := ErrorResponse{
		Error: msg,
	}
	return c.JSON(status, res)
}

func NotFound(c echo.Context, msg string) error {
	return Error(c, http.StatusNotFound, msg)
}

func InternalServerError(c echo.Context, msg string) error {
	return Error(c, http.StatusInternalServerError, msg)
}
