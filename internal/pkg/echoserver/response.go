package echoserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Meta struct {
	Page       int  `json:"page"`
	Size       int  `json:"size"`
	TotalCount int  `json:"totalCount"`
	TotalPage  int  `json:"totalPage"`
	HasNext    bool `json:"hasNext"`
}

type CommonResponse struct {
	Meta *Meta `json:"meta"`
	Data *any  `json:"data"`
}

func OK(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, CommonResponse{
		Data: &data,
	})
}
