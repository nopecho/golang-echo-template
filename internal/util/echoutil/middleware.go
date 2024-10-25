package echoutil

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type httpMethod string

type SkipRequest struct {
	method httpMethod
	path   string
}

const (
	get   httpMethod = "GET"
	post  httpMethod = "POST"
	patch httpMethod = "PATCH"
	put   httpMethod = "PUT"
	del   httpMethod = "DELETE"
)

func defaultMiddleware(e *echo.Echo) {
	e.Use(logMiddleware())
	e.Use(middleware.Recover())
	e.Use(jwtMiddleware())
}

func logMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper:    logSkipper(),
		LogMethod:  true,
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("method", v.Method).
				Str("uri", v.URI).
				Int("status", v.Status).
				Str("latency", v.Latency.String()).
				Msg("API")
			return nil
		},
	})
}

func logSkipper() middleware.Skipper {
	skipPaths := []*SkipRequest{
		{method: get, path: "/"},
		{method: get, path: "/health"},
		{method: get, path: "/favicon.ico"},
	}
	return func(c echo.Context) bool {
		method := c.Request().Method
		path := c.Path()
		for _, p := range skipPaths {
			if p.method == httpMethod(method) && p.path == path {
				return true
			}
		}
		return false
	}
}
