package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/noah-platform/noah/pkg/response"
	"github.com/rs/zerolog/log"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

const userIDHeaderName = "X-Noah-User-Id"
const userIDContextKey = "userID"

func (m *Middleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		l := log.With().Str("requestId", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()

		userID := c.Request().Header.Get(userIDHeaderName)
		if userID == "" {
			l.Warn().Str("userID", userID).Msg("[Auth.Middleware] no auth header found")

			return response.Unauthorized(c, "unauthorized")
		}

		c.Set(userIDContextKey, userID)

		l.Info().Str("userID", userID).Msg("[Auth.Middleware] authenticated")

		return next(c)
	}
}

func (m *Middleware) GetUserID(c echo.Context) string {
	userID, ok := c.Get(userIDContextKey).(string)
	if !ok {
		panic("[Auth.GetUserID] invalid userID in context")
	}

	return userID
}
