package echo

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/nopecho/golang-template/internal/util/common"
	"github.com/nopecho/golang-template/internal/util/jwtutil"
)

const (
	JwtContextKey = "jwtUser"
)

func jwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper:       jwtSkipper(),
		TokenLookup:   "header:Authorization:Bearer ",
		SigningKey:    []byte(common.GetEnv("JWT_SECRET", "secret")),
		NewClaimsFunc: jwtUserClaims(),
		ContextKey:    JwtContextKey,
		SigningMethod: jwtutil.AlgorithmHS512,
	})
}

func jwtSkipper() func(c echo.Context) bool {
	authSkipPaths := []*SkipRequest{
		{method: get, path: "/"},
		{method: get, path: "/health"},
		{method: get, path: "/favicon.ico"},
	}
	return func(c echo.Context) bool {
		method := c.Request().Method
		path := c.Path()
		for _, p := range authSkipPaths {
			if p.method == httpMethod(method) && p.path == path {
				return true
			}
		}
		return false
	}
}

func jwtUserClaims() func(c echo.Context) jwt.Claims {
	return func(c echo.Context) jwt.Claims {
		return &jwtutil.JwtUser{}
	}
}
