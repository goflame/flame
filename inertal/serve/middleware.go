package serve

import (
	"github.com/goflame/flame/inertal/response"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/middleware"
)

func handleMiddleware(ms *middleware.Middleware, rr *response.RootResponse, ctx *http.Context, s *Server) bool {
	for _, m := range *ms.GetHandlers() {
		e := m(ctx)

		if e != nil {
			err := e.GetError()
			if err != nil {
				s.handleError(ctx, err, e.GetStatus())
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
