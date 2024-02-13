package middleware

import "github.com/goflame/flame/pkg/handler"

func NewHandle(h ...handler.Handler) *Middleware {
	return New().UseMany(h...)
}
