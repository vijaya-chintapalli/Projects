package model

// RuleType is an enumeration for rule types
type RuleType int

const (
	RuleTypeStoreName RuleType = iota + 1
	RuleTypeItemMatch
	RuleTypeRetailerNameAlpha
	RuleTypeTotalRoundDollar
	RuleTypeTotalMultipleOfQuarter
	RuleTypeItemCountMultiplier
	RuleTypeItemDescriptionLength
	RuleTypePurchaseDayOdd
	RuleTypePurchaseTimeRange
)