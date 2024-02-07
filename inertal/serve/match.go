package serve

type Match struct {
	path string
}

func NewMatch(path string) *Match {
	return &Match{
		path: path,
	}
}

func (m *Match) Try(path string) (bool, map[string]string) {
	if path == m.path {
		return true, make(map[string]string)
	}
	return false, make(map[string]string)
}
