package auth

import (
	"errors"
	"github.com/goflame/flame/pkg/handler"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/response"
)

type Middleware struct {
	allow func(*http.Request) bool
}

func New(allow func(*http.Request) bool) handler.Handler {
	m := &Middleware{
		allow: allow,
	}
	return m.Middleware
}

func (m *Middleware) Middleware(res *http.Response, req *http.Request) *response.Err {
	if m.allow(req) {
		return res.Next()
	}
	return res.Error(errors.New("unauthorized")).Status(401)
}
