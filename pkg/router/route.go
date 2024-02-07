package router

import (
	"github.com/goflame/flame/pkg/handler"
	"github.com/goflame/flame/pkg/http/middleware"
)

type Route struct {
	Middlewares []*middleware.Middleware
	Method      string
	Path        string
	Handler     *handler.Handler
}

func (r *Route) Middleware(m *middleware.Middleware) {
	r.Middlewares = append(r.Middlewares, m)
}
