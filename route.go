package flame

func findRoute(name string, _ Map, f *Flame) string {
	// TODO: Add prop support
	routes := f.Router.Export()
	for _, route := range routes {
		if route.ExportName() == name {
			return route.Path
		}
	}
	return ""
}
