package flame

import (
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/response"
	"log"
	"testing"
)

func TestNewServer(t *testing.T) {
	app := New("FlameCore", true)

	app.Public(Root("/web"))
	app.Router.Get("/test", func(res *http.Response, req *http.Request) *response.Err {
		return res.String("Hello World not file!")
	})
	log.Fatal(app.Serve(8000))
}
