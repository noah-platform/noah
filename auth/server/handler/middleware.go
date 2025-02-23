package handler

import "github.com/labstack/echo/v4"

type AuthMiddleware interface {
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
	GetUserID(c echo.Context) string
}
