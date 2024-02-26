package flame

import (
	"fmt"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/middleware/auth"
	"github.com/goflame/flame/pkg/http/response"
	"github.com/goflame/flame/pkg/router"
	"log"
	"strings"
	"testing"
)

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
	}).Middleware(auth.New(func(req *http.Request) bool {
		return req.Query("token") == "secret"
	}))

	log.Fatal(app.Serve(8000))
}
