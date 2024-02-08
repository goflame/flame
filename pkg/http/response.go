package http

import (
	"encoding/json"
	rr "github.com/goflame/flame/inertal/response"
	"github.com/goflame/flame/pkg/http/response"
	nethttp "net/http"
)

type Response struct {
	rr      *rr.RootResponse
	code    int
	headers map[string]string
	next    bool
}

func NewResponse(rr *rr.RootResponse) *Response {
	return &Response{
		rr:      rr,
		code:    200,
		next:    false,
		headers: map[string]string{},
	}
}

func (r *Response) Status(code int) *Response {
	r.code = code
	return r
}

func (r *Response) Headers(h map[string]string) *Response {
	for k, v := range h {
		r.headers[k] = v
	}
	return r
}

func (r *Response) Header(name string, value string) *Response {
	r.headers[name] = value
	return r
}

func (r *Response) Next() *response.Err {
	r.rr.Next()
	return nil
}

func (r *Response) JSON(d interface{}) *response.Err {
	r.beforeSend()
	err := json.NewEncoder(*r.rr.ResponseWriter).Encode(d)
	if err != nil {
		return response.NewError(err)
	}
	return nil
}

func (r *Response) String(s string) *response.Err {
	r.beforeSend()
	rw := *r.rr.ResponseWriter
	_, err := rw.Write([]byte(s))
	if err != nil {
		return response.NewError(err)
	}
	return nil
}

func (r *Response) Text(s string) *response.Err {
	r.Headers(map[string]string{
		"Content-Type": "text/plain",
	})
	return r.String(s)
}

func (r *Response) Empty() *response.Err {
	return r.Text("")
}

func (r *Response) File(path string, req *Request) *response.Err {
	nethttp.ServeFile(*r.rr.ResponseWriter, req.NetRequest(), path)
	return nil
}

func (r *Response) Error(err error) *response.Err {
	return response.NewError(err)
}

func (r *Response) NetResponseWriter() nethttp.ResponseWriter {
	return *r.rr.ResponseWriter
}

func (r *Response) beforeSend() {
	for k, v := range r.headers {
		(*r.rr.ResponseWriter).Header().Set(k, v)
	}
}
