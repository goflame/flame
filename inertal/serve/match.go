package serve

type Match interface {
	Incoming(path string)
	TryPattern(pattern string) (bool, map[string]string)
}
