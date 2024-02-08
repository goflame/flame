package cors

import (
	"github.com/goflame/flame/pkg/handler"
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/response"
	nethttp "net/http"
	"strconv"
	"strings"
)

type AllowHandler func(*http.Response, *http.Request) bool

type Cors struct {
	config Config
}

type Config struct {
	AllowHandler     *AllowHandler
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    string
	MaxAge           int
}

func New(conf Config) handler.Handler {
	c := &Cors{
		config: conf,
	}
	return c.Middleware
}

func (c Cors) Middleware(res *http.Response, req *http.Request) *response.Err {
	if req.Method() == nethttp.MethodOptions {
		if c.config.AllowHandler != nil {
			h := *c.config.AllowHandler
			if h(res, req) {
				nr := req.NetRequest()
				// TODO: finish configuring cors middleware
				res.Headers(map[string]string{
					"Access-Control-Allow-Origin":   nr.Header.Get("Origin"),
					"Access-Control-Allow-Methods":  nr.Method,
					"Access-Control-Expose-Headers": c.config.ExposeHeaders,
				})
			}
		} else {
			if len(c.config.AllowOrigins) > 0 {
				res.Header("Access-Control-Allow-Origin", strings.Join(c.config.AllowOrigins, ","))
			}
			if len(c.config.AllowMethods) > 0 {
				res.Header("Access-Control-Allow-Methods", strings.Join(c.config.AllowMethods, ","))
			}
			if len(c.config.AllowHeaders) > 0 {
				res.Header("Access-Control-Allow-Headers", strings.Join(c.config.AllowHeaders, ","))
			}
			if c.config.AllowCredentials {
				res.Header("Access-Control-Allow-Credentials", "true")
			}
			if c.config.ExposeHeaders != "" {
				res.Header("Access-Control-Expose-Headers", c.config.ExposeHeaders)
			}
			if c.config.MaxAge > 0 {
				res.Header("Access-Control-Max-Age", strconv.Itoa(c.config.MaxAge))
			}
		}
		return nil
	}
	return res.Next()
}
