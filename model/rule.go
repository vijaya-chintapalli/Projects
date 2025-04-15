package model

type RuleType int

const (
	RetailerNameAlpha RuleType = iota + 1
	TotalRoundDollar
	TotalMultipleOfQuarter
	ItemCountMultiplier
	ItemDescriptionLength
	PurchaseDayOdd
	PurchaseTimeRange
)
