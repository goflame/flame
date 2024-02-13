package router

type RuleCheck func(string, []string) (string, bool)
