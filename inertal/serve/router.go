package serve

import (
	"errors"
	"github.com/goflame/flame/inertal/response"
	"github.com/goflame/flame/pkg/http"
)

type Router struct {
	server *Server
}

func NewRouter(s *Server) *Router {
	return &Router{
		server: s,
	}
}

func (r *Router) HandleRoute(m Match, rr *response.RootResponse, ctx *http.Context) {
	m.Incoming(ctx.Request.Net().URL.Path)
	for _, route := range r.server.Routes {
		if route.Method != ctx.Request.Method() {
			continue
		}

		ok, props := m.TryPattern(route.Path)

		if !ok {
			continue
		}

		ctx.Request.Props = props

		for _, mw := range route.Middlewares {
			if !handleMiddleware(mw, rr, ctx, r.server) {
				return
			}
		}

		h := *route.Handler
		err := h(ctx)

		if e := err.GetError(); e != nil {
			r.server.handleError(ctx, e, err.GetStatus())
			return
		}

		return
	}

	r.server.handleError(ctx, errors.New("this page could not be found"), 404)
}
