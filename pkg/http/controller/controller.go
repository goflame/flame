package controller

import (
	"fmt"
	"github.com/goflame/flame/pkg/handler"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/middleware"
	"github.com/goflame/flame/pkg/http/response"
	"reflect"
)

type Base struct {
	middlewares []*middleware.Middleware
}

func (c Base) AddMiddleware(m *middleware.Middleware) {
	c.middlewares = append(c.middlewares, m)
}

type IController interface {
	Init(c *http.Context) *response.Err
	AddMiddleware(m *middleware.Middleware)
}

func Adaptor(c IController, m string) handler.Handler {
	return func(ctx *http.Context) *response.Err {
		c.Init(ctx)
		method := reflect.ValueOf(c).MethodByName(m)

		if !method.IsValid() {
			return response.NewError(fmt.Sprintf("The method [%v] does not exist in the controller", m))
		}

		result := method.Call(nil)

		err := result[0].Interface()

		return err.(*response.Err)
	}
}
