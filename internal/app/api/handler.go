package api

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Group *echo.Group
}

func NewHandler(v *echo.Group) *Handler {
	return &Handler{
		Group: v,
	}
}
