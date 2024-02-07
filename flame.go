package flame

import (
	"fmt"
	"github.com/goflame/flame/inertal/serve"
	"github.com/goflame/flame/pkg/http/middleware"
	"github.com/goflame/flame/pkg/router"
	nethttp "net/http"
)

type Flame struct {
	appName    string
	envFile    string
	Middleware *middleware.Middleware
	Router     *router.Router
	Debug      bool
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

func (f *Flame) Serve(port int) error {
	fmt.Printf("Server listening on port %v\n", port)
	handler := serve.New(f.Router.Export(), f.Middleware, f.Debug)
	return nethttp.ListenAndServe(fmt.Sprintf(":%v", port), handler)
}
