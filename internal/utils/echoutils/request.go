package echoutils

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func DefaultQuery(c echo.Context, key string, defaultValue string) string {
	value := c.QueryParam(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func DefaultInt64Query(c echo.Context, key string, defaultValue int64) int64 {
	intValue := DefaultQuery(c, key, strconv.FormatInt(defaultValue, 10))
	value, err := strconv.ParseInt(intValue, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

func DefaultIntQuery(c echo.Context, key string, defaultValue int) int {
	intValue := DefaultQuery(c, key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(intValue)
	if err != nil {
		return defaultValue
	}
	return value
}
