package serve

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/goflame/flame/pkg/http/middleware"
	"github.com/goflame/flame/pkg/router"
	"golang.org/x/term"
	"math"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Middleware *middleware.Middleware
	Routes     []*router.Route
	Debug      bool
}

func New(r []*router.Route, m *middleware.Middleware, d bool) *Server {
	return &Server{
		Routes:     r,
		Middleware: m,
		Debug:      d,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.Debug == true {
		s.printRequest(r.Method, r.URL.Path)
	}
	NewRouter(s).HandleRoute(w, r)
}

func (s *Server) printRequest(m string, p string) {
	var method string

	switch m {
	case "GET":
		method = color.New(color.FgGreen).Sprintf("%v", m)
		break
	case "POST":
		method = color.New(color.FgBlue).Sprintf("%v", m)
		break
	case "PUT":
		method = color.New(color.FgCyan).Sprintf("%v", m)
		break
	case "PATCH":
		method = color.New(color.FgHiYellow).Sprintf("%v", m)
		break
	case "DELETE":
		method = color.New(color.FgRed).Sprintf("%v", m)
		break
	}

	w, _, err := term.GetSize(0)
	t := time.Now().UTC().Format(time.TimeOnly)
	l := w - len(fmt.Sprintf("[ %v ]%v %v", t, m, p)) - 1
	if err != nil || l <= 0 {
		fmt.Printf("[%v] %v", method, p)
	} else {
		dot := ". "
		dots := strings.Repeat(dot, int(math.Round(float64(l/2))))
		fmt.Printf("[ %v ] %v %v%v\n", method, color.New(color.FgHiBlack).Sprintf("%v", t), color.New(color.FgHiBlack).Sprintf("%v", dots), p)
	}
}
