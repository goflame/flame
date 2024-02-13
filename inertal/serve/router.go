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

func (r *Router) HandleRoute(m Match, rr *response.RootResponse, res *http.Response, req *http.Request) {
	m.Incoming(req.Net().URL.Path)
	for _, route := range r.server.Routes {
		if route.Method != req.Method() {
			continue
		}

		ok, props := m.TryPattern(route.Path)

		if !ok {
			continue
		}

		req.Props = props

		for _, mw := range route.Middlewares {
			if !handleMiddleware(mw, rr, res, req, r.server) {
				return
			}
		}

		h := *route.Handler
		err := h(*res, req)

		if e := err.GetError(); e != nil {
			r.server.handleError(res, req, e, err.GetStatus())
			return
		}

		return
	}

	r.server.handleError(res, req, errors.New("this page could not be found"), 404)
}
