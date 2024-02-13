package router

import (
	"github.com/goflame/flame/utils"
	"strconv"
)

type Rules map[string]RuleCheck

type defaultRules struct{}

func (defaultRules) int(value string, _ []string) (string, bool) {
	_, err := strconv.Atoi(value)
	return value, err == nil
}

func (defaultRules) in(value string, args []string) (string, bool) {
	in := false
	for _, arg := range args {
		if arg == value {
			in = true
			break
		}
	}
	return value, in
}

func (defaultRules) float64(value string, _ []string) (string, bool) {
	_, err := strconv.ParseFloat(value, 64)
	return value, err == nil
}

func (defaultRules) between(value string, args []string) (string, bool) {
	if len(args) != 2 {
		return "", false
	}
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "", false
	}
	minVal, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return "", false
	}
	maxVal, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return "", false
	}
	return value, val >= minVal && val <= maxVal
}

func (defaultRules) bool(value string, _ []string) (string, bool) {
	return value, value == "true" || value == "false"
}

func (d defaultRules) Make() Rules {
	return Rules{
		"int":     d.int,
		"in":      d.in,
		"float":   d.float64,
		"between": d.between,
		"bool":    d.bool,
	}
}

func (r Rules) Try(rule string, value string, args []string) (string, bool) {
	validator, ok := r[rule]

	if !ok {
		utils.Log.Warn("Rule ", rule, " does not exist")
		return "", false
	}

	val, suc := validator(value, args)
	return val, suc
}
