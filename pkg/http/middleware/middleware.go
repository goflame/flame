package middleware

import (
	"github.com/goflame/flame/pkg/handler"
	"net/http"
)

type Handler *http.Handler

type Middleware struct {
	handlers []handler.Handler
}

func New() *Middleware {
	return &Middleware{
		handlers: []handler.Handler{},
	}
}

func (m *Middleware) Use(h handler.Handler) *Middleware {
	m.handlers = append(m.handlers, h)
	return m
}

func (m *Middleware) UseMany(h ...handler.Handler) *Middleware {
	m.handlers = append(m.handlers, h...)
	return m
}

func (m *Middleware) GetHandlers() *[]handler.Handler {
	return &m.handlers
}
