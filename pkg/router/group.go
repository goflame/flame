package router

import (
	"github.com/goflame/flame/pkg/handler"
	"github.com/goflame/flame/pkg/http/middleware"
)

type Group struct {
	router      *Router
	prefix      string
	Middlewares []*middleware.Middleware
	routes      []*Route
}

func NewGroup(prefix string, router *Router) *Group {
	return &Group{
		router:      router,
		prefix:      prefix,
		Middlewares: []*middleware.Middleware{},
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
	g.Middlewares = append(g.Middlewares, m)
	return g
}

func (g *Group) Apply() {
	for _, r := range g.routes {
		route := g.router.addRoute(r.Method, r.Path, *r.Handler)
		for _, m := range g.Middlewares {
			route.Middleware(m)
		}
	}
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
	r := &Route{
		Method:  method,
		Path:    g.path(path),
		Handler: &handler,
	}
	g.routes = append(g.routes, r)
	return r
}
