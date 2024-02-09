package http

import (
	"github.com/goflame/flame/pkg/http/request"
	"net/http"
)

type Request struct {
	req   *http.Request
	Path  string
	Body  *request.BodyReader
	Props map[string]string
}

func NewRequest(req *http.Request) *Request {
	return &Request{
		req:  req,
		Path: req.URL.Path,
		Body: request.NewBodyReader(&req.Body),
	}
}

func (r *Request) Method() string {
	return r.req.Method
}

func (r *Request) Query(s string) string {
	return r.req.URL.Query().Get(s)
}

func (r *Request) Header(s string) string {
	return r.req.Header.Get(s)
}

func (r *Request) Net() *http.Request {
	return r.req
}
