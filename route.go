package flame

import (
	"fmt"
	"github.com/goflame/flame/pkg/router"
	"strings"
)

func findRoute(name string, m map[string]any, r *router.Router) string {
	routes := r.Export()
	for _, route := range routes {
		if route.ExportName() == name {
			rt := route.Path
			if strings.HasSuffix(rt, "/") {
				rt = rt[0 : len(rt)-1]
			}
			if strings.HasPrefix(rt, "/") {
				rt = rt[1:]
			}

			var built string
			sp := strings.Split(rt, "/")
			for _, part := range sp {
				if !strings.HasPrefix(part, "{") {
					built += "/" + part
				} else {
					part = part[1 : len(part)-1]
					if strings.Contains(part, "<") {
						sp = strings.Split(part, "<")
						built += fmt.Sprintf("/%s", m[sp[0]])
					} else {
						built += fmt.Sprintf("/%s", m[part])
					}
				}
			}
			return built
		}
	}
	return ""
}
