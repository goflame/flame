package flame

import (
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/middleware/auth"
	"github.com/goflame/flame/pkg/http/response"
	"github.com/goflame/flame/pkg/router"
	"log"
	"strings"
	"testing"
)

type Env struct {
	Name string `env:"NAME"`
}

func TestNewServer(t *testing.T) {
	app := New("FlameCore", true)

	app.PublicDir(Root("/web"))

	app.Router.Rule("jpeg", func(value string, _ []string) (string, bool) {
		if !strings.HasSuffix(value, ".jpeg") {
			return "", false
		}
		return value[:len(value)-len(".jpeg")], true
	})

	app.Router.Get("/test/{image<jpeg>}", func(res http.Response, req *http.Request) *response.Err {
		return res.JSON(Map{
			"image": req.Props["image"],
		})
	})

	app.Router.Group("/http", func(g *router.Group) {
		g.Get("/", func(res http.Response, req *http.Request) *response.Err {
			return res.Text("Hello world")
		})
	}).Middleware(auth.New(func(req *http.Request) bool {
		return req.Query("token") == "secret"
	}))

	app.Router.Get("/", func(res http.Response, req *http.Request) *response.Err {
		return res.JSON(Map{
			"id":  5,
			"url": req.Net().URL.String(),
		})
	})

	log.Fatal(app.Serve(8000))
}
