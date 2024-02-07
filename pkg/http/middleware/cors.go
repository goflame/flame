package middleware

import (
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/response"
	nethttp "net/http"
)

func Cors(res *http.Response, req *http.Request) *response.Err {
	if req.Method() == nethttp.MethodOptions {
		return res.Headers(map[string]string{
			"Content-Type": "text/plain",
		}).Empty()
	}
	return res.Next()
}
