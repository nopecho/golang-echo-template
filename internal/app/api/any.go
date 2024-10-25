package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nopecho/golang-template/internal/app/domain"
	echoutil "github.com/nopecho/golang-template/internal/util/echoutil"
)

type AnyRouter struct {
	svc *domain.AnyService
}

func NewAnyRouter(service *domain.AnyService) *AnyRouter {
	return &AnyRouter{
		svc: service,
	}
}

func (r *AnyRouter) Route(g *echo.Group) {
	g.GET("/domain/:id", r.get)
	g.POST("/domain", r.create)
	g.PATCH("/domain/:id", r.update)
}

func (r *AnyRouter) get(c echo.Context) error {
	var param GetParams
	if err := c.Bind(&param); err != nil {
		return echoutil.BadRequest(c, "Bad Request")
	}

	data, err := r.svc.GetById(param.ID)
	if err != nil {
		return echoutil.NotFound(c, "Not Found")
	}
	return echoutil.OK(c, data)
}

func (r *AnyRouter) create(c echo.Context) error {
	var body CreateRequest
	if err := c.Bind(&body); err != nil {
		return echoutil.BadRequest(c, "Bad Request")
	}

	domain, _ := r.svc.Create(&domain.AnyCreateCommand{
		Name: body.Name,
	})
	return echoutil.OK(c, domain)
}

func (r *AnyRouter) update(c echo.Context) error {
	var body UpdateRequest
	if err := c.Bind(&body); err != nil {
		return echoutil.BadRequest(c, "Bad Request")
	}

	updated, err := r.svc.Update(&domain.AnyUpdateCommand{
		ID:   body.ID,
		Name: body.Name,
	})
	if err != nil {
		return echoutil.NotFound(c, "Not Found")
	}

	return echoutil.OK(c, updated)
}
