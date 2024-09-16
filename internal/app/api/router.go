package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

// Router is an interface for echo routing
type Router interface {
	Route(e *echo.Echo)
}

// GroupRouter is an interface for echo versioned routing
type GroupRouter interface {
	route(e *echo.Group)
}

// Handler is a struct for handling echo routing with versioning
// Version is the version of the API prefix (e.g. v1 -> route: /api/v1/~)
type Handler struct {
	Version string
	Routers []GroupRouter
}

const prefix = "/api"

func NewHandler(version string) *Handler {
	return &Handler{
		Version: version,
	}
}

func (h *Handler) Register(gr ...GroupRouter) {
	h.Routers = append(h.Routers, gr...)
}

func (h *Handler) Route(e *echo.Echo) {
	group := e.Group(h.versioning())
	for _, vr := range h.Routers {
		vr.route(group)
	}
}

func (h *Handler) versioning() string {
	if h.Version == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", prefix, h.Version)
}
