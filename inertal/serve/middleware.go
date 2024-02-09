package serve

import (
	"github.com/goflame/flame/inertal/response"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/middleware"
)

func handleMiddleware(ms *middleware.Middleware, rr *response.RootResponse, res *http.Response, req *http.Request, s *Server) bool {
	for _, m := range *ms.GetHandlers() {
		e := m(res, req)

		if e != nil {
			err := e.GetError()
			if err != nil {
				s.handleError(res, req, err, e.GetStatus())
				return false
			}
		}

		cont := rr.CanContinue()
		rr.Reset()

		if !cont {
			return false
		}

		continue
	}
	return true
}
