package model

type RuleType string

const (
    RetailerNameAlpha       RuleType = "RetailerNameAlpha"
    TotalRoundDollar        RuleType = "TotalRoundDollar"
    TotalMultipleOfQuarter  RuleType = "TotalMultipleOfQuarter"
    ItemCountMultiplier     RuleType = "ItemCountMultiplier"
    TotalRoundDollar        RuleType = "TotalRoundDollar"
    TotalMultipleOfQuarter  RuleType = "TotalMultipleOfQuarter"
    ItemCountMultiplier     RuleType = "ItemCountMultiplier"
    ItemDescriptionLength   RuleType = "ItemDescriptionLength"
    PurchaseDayOdd          RuleType = "PurchaseDayOdd"
    PurchaseTimeRange       RuleType = "PurchaseTimeRange"
)
