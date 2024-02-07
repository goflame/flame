package serve

import (
	"fmt"
	"github.com/goflame/flame/inertal/response"
	"github.com/goflame/flame/pkg/http"
	nethttp "net/http"
)

type Router struct {
	server *Server
}

func NewRouter(s *Server) *Router {
	return &Router{
		server: s,
	}
}

func (m *Router) HandleRoute(w nethttp.ResponseWriter, r *nethttp.Request) {
	rr, res, req := m.convertRequest(w, r)
	for _, mw := range *m.server.Middleware.GetHandlers() {
		err := mw(res, req)
		if e := err.GetError(); e != nil {
			m.handleError(e, err.GetStatus())
			return
		}
		canContinue := rr.CanContinue()
		rr.Reset()
		if canContinue {
			continue
		}
		return
	}
	router := NewMatch(r.URL.Path)
	for _, route := range m.server.Routes {
		if route.Method != r.Method {
			continue
		}
		ok, props := router.Try(route.Path)
		if !ok {
			continue
		}
		req.Props = props
		for _, mw := range route.Middlewares {
			for _, mh := range *mw.GetHandlers() {
				err := mh(res, req)
				if e := err.GetError(); e != nil {
					m.handleError(e, err.GetStatus())
					return
				}
				canContinue := rr.CanContinue()
				rr.Reset()
				if canContinue {
					continue
				}
				return
			}
		}
		h := *route.Handler
		err := h(res, req)
		if e := err.GetError(); e != nil {
			m.handleError(e, err.GetStatus())
			return
		}
		return
	}
}

func (m *Router) convertRequest(w nethttp.ResponseWriter, r *nethttp.Request) (*response.RootResponse, *http.Response, *http.Request) {
	rr := response.NewRootResponse(&w)
	return rr, http.NewResponse(rr), http.NewRequest(r)
}

func (m *Router) handleError(err error, code int) {
	fmt.Println(err, code)
}
