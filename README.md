# Flame, a simple web framework

This framework built on top of the `net/http` package.

## Demo

### Creating a simple server

```go
package main

import (
	"github.com/goflame/flame"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/response"
	"log"
)

func main() {
	app := flame.New("Your app's name", true)

	app.Router.Get("/", func(res http.Response, _ *http.Request) *response.Err {
        return res.String("Hello world!")
    })

	log.Fatal(app.Serve(3000))
}
```