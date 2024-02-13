package serve

import (
	"github.com/goflame/flame/inertal/dev"
	"github.com/goflame/flame/inertal/response"
	"github.com/goflame/flame/pkg/config"
	"github.com/goflame/flame/pkg/handler"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/middleware"
	"github.com/goflame/flame/pkg/router"
	nethttp "net/http"
	"os"
	"strings"
)

type Server struct {
	wwwRoot    string
	Middleware *middleware.Middleware
	Routes     []*router.Route
	Debug      bool
	Eh         handler.ErrorHandler
	AppConfig  *config.App
}

func New(c *config.App, static string, r []*router.Route, eh handler.ErrorHandler, m *middleware.Middleware, d bool) *Server {
	return &Server{
		wwwRoot:    static,
		Routes:     r,
		Middleware: m,
		Debug:      d,
		Eh:         eh,
		AppConfig:  c,
	}
}

func (s *Server) ServeHTTP(w nethttp.ResponseWriter, r *nethttp.Request) {
	dr := dev.Request{}
	if s.Debug == true {
		dr.Log(r.Method, r.URL.Path)
	}

	rr, res, req := s.convertRequest(w, r)

	if !handleMiddleware(s.Middleware, rr, res, req, s) {
		return
	}

	if _, err := os.Stat(s.wwwRoot + r.URL.Path); err == nil {
		if !strings.HasSuffix(r.URL.Path, "/") {
			dr.FileLog(r.URL.Path)
			nethttp.ServeFile(w, r, s.wwwRoot+r.URL.Path)
			return
		}
	}

	NewRouter(s).HandleRoute(s.AppConfig.RouterMatch, rr, res, req)
}

func (*Server) convertRequest(w nethttp.ResponseWriter, r *nethttp.Request) (*response.RootResponse, *http.Response, *http.Request) {
	rr := response.NewRootResponse(&w)
	return rr, http.NewResponse(rr), http.NewRequest(r)
}

func (s *Server) handleError(res *http.Response, req *http.Request, err error, code int) {
	s.Eh(*res, req, err, code)
}
