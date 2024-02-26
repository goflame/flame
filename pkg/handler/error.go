package handler

import (
	"github.com/goflame/flame/pkg/http"
)

type ErrorHandler func(*http.Context, error, int)

func DefaultErrorHandler(c *http.Context, err error, status int) {
	_ = c.Status(status).
		JSON(map[string]interface{}{
			"error": err.Error(),
		})
}
