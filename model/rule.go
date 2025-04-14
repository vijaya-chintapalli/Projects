package model

type RuleType int

const (
	RuleTypeStoreName RuleType = iota + 1
	RuleTypeItemMatch
)
