package http

import "net/http"

type Handler interface {
	CanHandle(r *http.Request) bool
	Handle(w http.ResponseWriter, r *http.Request)
}

type CompositeHandler struct {
	handlers []Handler
}

func NewCompositeHandler(handlers ...Handler) *CompositeHandler {
	return &CompositeHandler{handlers: handlers}
}

func (c *CompositeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handler := range c.handlers {
		if handler.CanHandle(r) {
			handler.Handle(w, r)
			return
		}
	}
	http.NotFound(w, r)
}
