package router

import (
	"github.com/goflame/flame/pkg/handler"
	"github.com/goflame/flame/pkg/http/middleware"
)

type Group struct {
	router *Router
	prefix string
	routes []*Route
}

func NewGroup(prefix string, router *Router) *Group {
	return &Group{
		router: router,
		prefix: prefix,
	}
}

func (g *Group) Get(path string, handler handler.Handler) *Route {
	return g.route("GET", path, handler)
}

func (g *Group) Post(path string, handler handler.Handler) *Route {
	return g.route("POST", path, handler)
}

func (g *Group) Put(path string, handler handler.Handler) *Route {
	return g.route("PUT", path, handler)
}

func (g *Group) Patch(path string, handler handler.Handler) *Route {
	return g.route("PATCH", path, handler)
}

func (g *Group) Delete(path string, handler handler.Handler) *Route {
	return g.route("DELETE", path, handler)
}

func (g *Group) Middleware(m *middleware.Middleware) *Group {
	for _, route := range g.routes {
		route.Middleware(m)
	}
	return g
}

func (g *Group) path(path string) string {
	if string(g.prefix[len(g.prefix)-1]) != "/" {
		if string(path[0]) == "/" {
			return g.prefix + path
		}
		return g.prefix + "/" + path
	}
	if string(path[0]) == "/" {
		return g.prefix + path[1:]
	}
	return g.prefix + path
}

func (g *Group) route(method string, path string, handler handler.Handler) *Route {
	r := g.router.addRoute(method, g.path(path), handler)
	g.routes = append(g.routes, r)
	return r
}
