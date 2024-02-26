package auth

import (
	"github.com/goflame/flame/pkg/handler"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/middleware"
	"github.com/goflame/flame/pkg/http/response"
)

type Middleware struct {
	allow func(*http.Request) bool
}

func New(allow func(*http.Request) bool) *middleware.Middleware {
	m := &Middleware{
		allow: allow,
	}
	return middleware.New().Use(m.Middleware)
}

func NewHandle(allow func(*http.Request) bool) handler.Handler {
	m := &Middleware{
		allow: allow,
	}
	return m.Middleware
}

func (m *Middleware) Middleware(c *http.Context) *response.Err {
	if m.allow(c.Request) {
		c.Set("auth", true)
		return c.Next()
	}
	return c.Error("unauthorized").Status(401)
}
