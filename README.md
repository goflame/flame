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
	app := flame.New(true)

	app.Router.Get("/", func(c *http.Context) *response.Err {
		return c.String("Hello world!")
	})

	log.Fatal(app.Serve(3000))
}
```

### Default routing rules

```go
func main() {
    app.Router.Get("/numbers/count<int>", func(c *http.Context) *response.Err {
        return c.String(c.Prop("count"))
    })
}
```


| Name      | Second Header                                                                                                         |
|-----------|-----------------------------------------------------------------------------------------------------------------------|
| `int`     | Checks if the parameter is an integer or not.                                                                         |
| `in`      | Use with multiple args, like this: `/fruit<in[apple;banana]>`. This checks if the value is in the list or not.        |
| `float`   | Checks if the parameter is an float or not.                                                                           |
| `between` | Use with 2 integers, like this: `/something<in[1;5]>`. This checks if the value is an integer between the two number. |
| `bool`    | Bool checks if it's a boolean or not. Works with both `true` `false` and `0` `1`                                      |


### Advanced routing with custom rules and props

```go

func main() {
    app.Router.Rule("image", func(value string, args []string) (string, bool) {
        allowedExtensions := append([]string{"jpeg", "png", "jpg", "webp"}, args...)
        for _, ext := range allowedExtensions {
            if strings.HasSuffix(value, "."+ext) {
            return value[:len(value)-len("."+ext)], true
            }
        }
        return "", false
    })	

    app.Router.Get("/profile_images/profileImageId<image[gif]>", func(c *http.Context) *response.Err {
        return c.JSON(flame.Map{
            "image_id": c.Prop("profileImageId"),
        })
    })
}
```