package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/nopecho/golang-template/internal/app/usecase"
	"github.com/nopecho/golang-template/pkg/echoserver"
)

type DomainHandler struct {
	Group   *echo.Group
	usecase *DomainUsecase
}

type DomainUsecase struct {
	Loader usecase.Loader
}

func NewDomainHandler(v *echo.Group, usecase *DomainUsecase) *DomainHandler {
	return &DomainHandler{
		Group:   v.Group("/domain"),
		usecase: usecase,
	}
}

func (h *DomainHandler) Routing() {
	h.Group.GET("", h.get)
	h.Group.POST("", h.create)
}

func (h *DomainHandler) get(c echo.Context) error {
	return c.JSON(200, echoserver.Map{
		"get": "domain",
	})
}

func (h *DomainHandler) create(c echo.Context) error {
	data, err := h.usecase.Loader.Load()
	if err != nil {
		return err
	}
	return c.JSON(200, data)
}
