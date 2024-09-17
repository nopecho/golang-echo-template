package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nopecho/golang-template/internal/app/svc"
	echoutil "github.com/nopecho/golang-template/internal/util/echo"
)

type DomainRouter struct {
	svc *svc.DomainService
}

func NewDomainRouter(service *svc.DomainService) *DomainRouter {
	return &DomainRouter{
		svc: service,
	}
}

func (h *DomainRouter) route(g *echo.Group) {
	g.GET("/domain/:id", h.get)
	g.POST("/domain", h.create)
	g.PATCH("/domain/:id", h.update)
}

func (h *DomainRouter) get(c echo.Context) error {
	var param GetParams
	if err := c.Bind(&param); err != nil {
		return echoutil.BadRequest(c, "Bad Request")
	}

	data, err := h.svc.GetById(param.ID)
	if err != nil {
		return echoutil.NotFound(c, "Not Found")
	}
	return echoutil.OK(c, data)
}

func (h *DomainRouter) create(c echo.Context) error {
	var body CreateRequest
	if err := c.Bind(&body); err != nil {
		return echoutil.BadRequest(c, "Bad Request")
	}

	domain, _ := h.svc.Create(&svc.DomainCreateCommand{
		Name: body.Name,
	})
	return echoutil.OK(c, domain)
}

func (h *DomainRouter) update(c echo.Context) error {
	var body UpdateRequest
	if err := c.Bind(&body); err != nil {
		return echoutil.BadRequest(c, "Bad Request")
	}

	updated, err := h.svc.Update(&svc.DomainUpdateCommand{
		ID:   body.ID,
		Name: body.Name,
	})
	if err != nil {
		return echoutil.NotFound(c, "Not Found")
	}

	return echoutil.OK(c, updated)
}
