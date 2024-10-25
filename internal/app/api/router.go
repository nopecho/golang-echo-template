package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

// EchoRouter is an interface for echo routing
type EchoRouter interface {
	Routing(e *echo.Echo)
}

// Router is an interface for echo versioned routing
type Router interface {
	Route(e *echo.Group)
}

// Handler is a struct for handling echo routing with versioning
// Version is the version of the API prefix (e.g. v1 -> Route: /api/v1/~)
type Handler struct {
	Version string
	Routers []Router
}

const prefix = "/api"

func NewVersionHandler(version string) *Handler {
	return &Handler{
		Version: version,
	}
}

func NewRootHandler() *Handler {
	return NewVersionHandler("")
}

func (h *Handler) Register(r ...Router) {
	h.Routers = append(h.Routers, r...)
}

func (h *Handler) Routing(e *echo.Echo) {
	group := e.Group(h.versioning())
	for _, r := range h.Routers {
		r.Route(group)
	}
}

func (h *Handler) versioning() string {
	if h.Version == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", prefix, h.Version)
}
