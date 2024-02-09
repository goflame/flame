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
	name        string
}

func (r *Route) Name(s string) *Route {
	r.name = s
	return r
}

func (r *Route) ExportName() string {
	return r.name
}

func (r *Route) Middleware(m *middleware.Middleware) {
	r.Middlewares = append(r.Middlewares, m)
}
