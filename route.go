package flame

import (
	"strings"
	"fmt"
)

func findRoute(name string, m Map, f *Flame) string {
	// FIXIT: 2 Slashes at start
	routes := f.Router.Export()
	for _, route := range routes {
		if route.ExportName() == name {
			var built string
			sp := strings.Split(route.Path, "/")
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
