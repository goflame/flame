package serve

type Match struct {
	path string
}

func NewMatch(path string) *Match {
	return &Match{
		path: path,
	}
}

func (m *Match) Try(incomingPath string) (bool, map[string]string) {
	if incomingPath == m.path {
		return true, make(map[string]string)
	}
	return false, make(map[string]string)
}
