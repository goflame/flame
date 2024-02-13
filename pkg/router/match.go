package router

import (
	"github.com/goflame/flame/utils"
	"strings"
)

type Match struct {
	path  string
	rules *Rules
}

func (m *Match) SetRules(rules *Rules) {
	m.rules = rules
}

func (m *Match) Incoming(path string) {
	m.path = m.removeUnnecessarySlashes(path)
}

func (m *Match) TryPattern(pattern string) (bool, map[string]string) {
	props := make(map[string]string)

	pattern = m.removeUnnecessarySlashes(pattern)

	if m.path == pattern {
		return true, nil
	}

	params := strings.Split(pattern, "/")
	routeParams := strings.Split(m.path, "/")

	if len(params) != len(routeParams) {
		return false, nil
	}

	for i, p := range params {
		if !strings.HasPrefix(p, "{") && !strings.HasSuffix(p, "}") {
			if p != routeParams[i] {
				return false, nil
			}
		}

		name, val, ok := m.validatePropRule(p[1:len(p)-1], routeParams[1])
		if !ok {
			return false, nil
		}

		props[name] = val
	}

	return true, props
}

func (*Match) removeUnnecessarySlashes(s string) string {
	if strings.HasPrefix(s, "/") {
		s = s[1:]
	}
	if strings.HasSuffix(s, "/") {
		s = s[:len(s)-1]
	}
	return s
}

func (m *Match) validatePropRule(p string, v string) (string, string, bool) {
	if !strings.Contains(p, "<") || !strings.Contains(p, ">") {
		return p, v, true
	}

	split := strings.Split(p, "<")
	if len(split) != 2 {
		return p, "", false
	}

	name := split[0]
	rule := split[1][:len(split[1])-1]
	var args []string

	if strings.Contains(rule, "[") && strings.Contains(rule, "]") {
		ruleSplit := strings.Split(rule, "[")

		if len(ruleSplit) < 2 {
			utils.Log.Warn("No rule parameter found")
		} else {
			rule = ruleSplit[0]
			args = strings.Split(ruleSplit[1][:len(ruleSplit[1])-1], ";")
		}
	}

	if val, suc := m.rules.Try(rule, v, args); suc {
		return name, val, true
	}

	return "", "", false
}
