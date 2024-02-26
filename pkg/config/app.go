package config

import (
	"github.com/goflame/flame/pkg/router"
)

type Match interface {
	SetRules(*router.Rules)
	Incoming(string)
	TryPattern(string) (bool, map[string]string)
}

type App struct {
	RouterMatch Match
}
