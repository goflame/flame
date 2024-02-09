package handler

import "github.com/goflame/flame/pkg/http"

type ErrorHandler func(*http.Response, *http.Request, error, int)

func DefaultErrorHandler(res *http.Response, _ *http.Request, err error, status int) {
	_ = res.Status(status).
		JSON(map[string]interface{}{
			"error": err.Error(),
		})
}
