package http

import (
	"github.com/goflame/flame/pkg/http/request"
	"net/http"
)

type Request struct {
	req   *http.Request
	Path  string
	Body  request.BodyReader
	Props map[string]string
}

func NewRequest(req *http.Request) *Request {
	return &Request{
		req:  req,
		Path: req.URL.Path,
	}
}

func (r *Request) Method() string {
	return r.req.Method
}

func (r *Request) NetRequest() *http.Request {
	return r.req
}
