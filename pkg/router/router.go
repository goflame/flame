package router

import (
	"github.com/goflame/flame/pkg/handler"
)

type Router struct {
	routes []*Route
}

func New() *Router {
	return &Router{
		routes: []*Route{},
	}
}

func (r *Router) Get(path string, handler handler.Handler) *Route {
	return r.addRoute("GET", path, handler)
}

func (r *Router) Post(path string, handler handler.Handler) *Route {
	return r.addRoute("POST", path, handler)
}

func (r *Router) Put(path string, handler handler.Handler) *Route {
	return r.addRoute("PUT", path, handler)
}

func (r *Router) Patch(path string, handler handler.Handler) *Route {
	return r.addRoute("PATCH", path, handler)
}

func (r *Router) Delete(path string, handler handler.Handler) *Route {
	return r.addRoute("DELETE", path, handler)
}

func (r *Router) Group(prefix string, f func(*Group)) *Group {
	g := NewGroup(prefix, r)
	f(g)
	return g
}

func (r *Router) Export() []*Route {
	return r.routes
}

func (r *Router) addRoute(method string, path string, handler handler.Handler) *Route {
	route := &Route{
		Method:  method,
		Path:    path,
		Handler: &handler,
	}
	r.routes = append(r.routes, route)
	return route
}
