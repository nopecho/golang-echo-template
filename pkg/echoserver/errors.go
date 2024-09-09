package echoserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	Code    any    `json:"code"`
	Message string `json:"message"`
	Data    *any   `json:"data"`
}

func Error(c echo.Context, code int, message string) error {
	return c.JSON(code, ErrorResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func BadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    nil,
	})
}

func NotFound(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, ErrorResponse{
		Code:    http.StatusNotFound,
		Message: message,
		Data:    nil,
	})
}
