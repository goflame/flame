package response

import "net/http"

type RootResponse struct {
	ResponseWriter *http.ResponseWriter
	next           bool
}

func NewRootResponse(rw *http.ResponseWriter) *RootResponse {
	return &RootResponse{
		ResponseWriter: rw,
	}
}

func (r *RootResponse) Next() {
	r.next = true
}

func (r *RootResponse) Reset() {
	r.next = false
}

func (r *RootResponse) CanContinue() bool {
	return r.next
}
