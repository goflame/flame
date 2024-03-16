package flame

import (
	"fmt"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/controller"
	"github.com/goflame/flame/pkg/http/middleware/auth"
	"github.com/goflame/flame/pkg/http/response"
	"github.com/goflame/flame/pkg/router"
	"log"
	"strings"
	"testing"
)

type TestController struct {
	controller.Base
	ctx *http.Context
}

func (t *TestController) Init(c *http.Context) *response.Err {
	t.ctx = c
	return c.Next()
}

func (t *TestController) Hello() *response.Err {
	return t.ctx.JSON(Map{
		"hello": "world",
	})
}

func TestNewServer(t *testing.T) {
	app := New(true)

	app.PublicDir(Root("/web"))

	app.Router.Rule("image", func(value string, args []string) (string, bool) {
		allowedExtensions := append([]string{"jpeg", "png", "jpg", "webp"}, args...)

		for _, ext := range allowedExtensions {
			if strings.HasSuffix(value, "."+ext) {
				return value[:len(value)-len("."+ext)], true
			}
		}
		return "", false
	})

	app.Router.Get("/test/{profileImage<image[gif]>}", func(c *http.Context) *response.Err {
		return c.JSON(Map{
			"image": c.Prop("profileImage"),
		})
	}).Name("image")

	app.Router.Group("/http", func(g *router.Group) {
		g.Get("/", func(c *http.Context) *response.Err {
			fmt.Println(c.Get("auth"))
			return c.String(app.Route("image", Map{
				"image": "goofy.jpeg",
			}))
		})

		g.Get("/gt/{id}", func(c *http.Context) *response.Err {
			return c.Text(c.Prop("id"))
		})
	}).Middleware(auth.New(func(req *http.Request) bool {
		return req.Query("token") == "secret"
	}))

	app.Router.Get("/controller", controller.Adaptor(&TestController{}, "Hello"))

	log.Fatal(app.Serve(8000))
}
