package flame

import (
	"fmt"
	"github.com/goflame/flame/inertal/console"
	"github.com/goflame/flame/inertal/serve"
	"github.com/goflame/flame/pkg/http/middleware"
	"github.com/goflame/flame/pkg/router"
	nethttp "net/http"
	"strings"
)

type Flame struct {
	appName    string
	envFile    string
	Middleware *middleware.Middleware
	Router     *router.Router
	Debug      bool
	wwwRoot    string
}

func New(name string, debug bool) *Flame {
	return &Flame{
		appName:    name,
		envFile:    "",
		Router:     router.New(),
		Middleware: middleware.New(),
		Debug:      debug,
	}
}

func (f *Flame) DotEnv(path string) {
	f.envFile = path
}

func (f *Flame) PublicDir(path string) *Flame {
	if strings.HasSuffix(path, "/") {
		f.wwwRoot = path[0 : len(path)-2]
	} else {
		f.wwwRoot = path
	}
	return f
}

func (f *Flame) Serve(port int) error {
	console.NewInfoPrint().Listen(port)
	handler := serve.New(f.wwwRoot, f.Router.Export(), f.Middleware, f.Debug)
	return nethttp.ListenAndServe(fmt.Sprintf(":%v", port), handler)
}
