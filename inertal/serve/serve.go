package serve

import (
	"github.com/goflame/flame/inertal/dev"
	"github.com/goflame/flame/pkg/http/middleware"
	"github.com/goflame/flame/pkg/router"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	wwwRoot    string
	Middleware *middleware.Middleware
	Routes     []*router.Route
	Debug      bool
}

func New(static string, r []*router.Route, m *middleware.Middleware, d bool) *Server {
	return &Server{
		wwwRoot:    static,
		Routes:     r,
		Middleware: m,
		Debug:      d,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dr := dev.Request{}
	if s.Debug == true {
		dr.Log(r.Method, r.URL.Path)
	}

	if _, err := os.Stat(s.wwwRoot + r.URL.Path); err == nil {
		if !strings.HasSuffix(r.URL.Path, "/") {
			dr.FileLog(r.URL.Path)
			http.ServeFile(w, r, s.wwwRoot+r.URL.Path)
			return
		}
	}

	NewRouter(s).HandleRoute(w, r)
}
