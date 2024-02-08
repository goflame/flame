package flame

import (
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/response"
	"log"
	"testing"
)

func TestNewServer(t *testing.T) {
	app := New("FlameCore", true)

	app.PublicDir(Root("/web"))
	app.Router.Get("/", func(res *http.Response, req *http.Request) *response.Err {
		return res.String("Hello")
	})
	log.Fatal(app.Serve(8000))
}
