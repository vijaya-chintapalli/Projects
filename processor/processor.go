package processor

import (
	"encoding/json"
	"fmt"
	"receipt-rule-engine-challenge/model"
)

// ruleProcessor is responsible for storing and processing rules
type ruleProcessor struct {
	rules map[model.RuleType]model.Rule
}

// NewProcessor initializes a new rule processor
func NewProcessor() *ruleProcessor {
	return &ruleProcessor{
		rules: make(map[model.RuleType]model.Rule),
	}
}

// AddRule adds a rule to the processor
func (rp *ruleProcessor) AddRule(ruleType model.RuleType, ruleDefinition string) error {
	var rule model.Rule
	var err error

	// Use the constants directly here
	switch ruleType {
	case model.RuleTypeStoreName:
		var retailerNameAlphaRule retailerNameAlphaRule
		err = json.Unmarshal([]byte(ruleDefinition), &retailerNameAlphaRule)
		if err == nil {
			rule = &retailerNameAlphaRule
		}
	case model.RuleTypeItemMatch:
		var totalRoundDollarRule totalRoundDollarRule
		err = json.Unmarshal([]byte(ruleDefinition), &totalRoundDollarRule)
		if err == nil {
			rule = &totalRoundDollarRule
		}
	case model.RuleTypeRetailerNameAlpha:
		var totalMultipleOfQuarterRule totalMultipleOfQuarterRule
		err = json.Unmarshal([]byte(ruleDefinition), &totalMultipleOfQuarterRule)
		if err == nil {
			rule = &totalMultipleOfQuarterRule
		}
	case model.RuleTypeTotalRoundDollar:
		var itemCountMultiplierRule itemCountMultiplierRule
		err = json.Unmarshal([]byte(ruleDefinition), &itemCountMultiplierRule)
		if err == nil {
			rule = &itemCountMultiplierRule
		}
	case model.RuleTypeTotalMultipleOfQuarter:
		var itemDescriptionLengthRule itemDescriptionLengthRule
		err = json.Unmarshal([]byte(ruleDefinition), &itemDescriptionLengthRule)
		if err == nil {
			rule = &itemDescriptionLengthRule
		}
	case model.RuleTypeItemCountMultiplier:
		var purchaseDayOddRule purchaseDayOddRule
		err = json.Unmarshal([]byte(ruleDefinition), &purchaseDayOddRule)
		if err == nil {
			rule = &purchaseDayOddRule
		}
	case model.RuleTypeItemDescriptionLength:
		var purchaseTimeRangeRule purchaseTimeRangeRule
		err = json.Unmarshal([]byte(ruleDefinition), &purchaseTimeRangeRule)
		if err == nil {
			rule = &purchaseTimeRangeRule
		}
	default:
		return fmt.Errorf("unsupported rule type: %v", ruleType)
	}

	if err != nil {
		return fmt.Errorf("failed to parse rule definition: %v", err)
	}

	rp.rules[ruleType] = rule
	return nil
}

// Process applies all the rules to the receipt and returns the total points
func (rp *ruleProcessor) Process(receipt model.Receipt) (int, error) {
	totalPoints := 0

	// Iterate through each rule and apply it to the receipt
	for _, rule := range rp.rules {
		points, err := rule.Evaluate(receipt)
		if err != nil {
			return 0, fmt.Errorf("error evaluating rule: %v", err)
		}
		totalPoints += points
	}

	return totalPoints, nil
}

// retailerNameAlphaRule checks if the retailer name consists of alphabetic characters
type retailerNameAlphaRule struct {
	PointsPerChar int `json:"pointsPerChar"`
}

func (r *retailerNameAlphaRule) Evaluate(receipt model.Receipt) (int, error) {
	count := 0
	for _, c := range receipt.StoreName {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			count++
		}
	}
	return count * r.PointsPerChar, nil
}

// totalRoundDollarRule checks if the total is a round dollar value
type totalRoundDollarRule struct {
	Points int `json:"points"`
}

func (r *totalRoundDollarRule) Evaluate(receipt model.Receipt) (int, error) {
	if receipt.Total == float64(int(receipt.Total)) {
		return r.Points, nil
	}
	return 0, nil
}

// totalMultipleOfQuarterRule checks if the total is a multiple of 0.25
type totalMultipleOfQuarterRule struct {
	Points int `json:"points"`
}

func (r *totalMultipleOfQuarterRule) Evaluate(receipt model.Receipt) (int, error) {
	if float64(int(receipt.Total*4)) == receipt.Total*4 {
		return r.Points, nil
	}
	return 0, nil
}

// itemCountMultiplierRule multiplies points by the number of items in the receipt
type itemCountMultiplierRule struct {
	Multiplier float64 `json:"multiplier"`
}

func (r *itemCountMultiplierRule) Evaluate(receipt model.Receipt) (int, error) {
	itemCount := len(receipt.Items)
	points := int(float64(itemCount) * r.Multiplier)
	return points, nil
}

// itemDescriptionLengthRule checks the length of the item descriptions
type itemDescriptionLengthRule struct {
	MaxLength int `json:"maxLength"`
	Points    int `json:"points"`
}

func (r *itemDescriptionLengthRule) Evaluate(receipt model.Receipt) (int, error) {
	for _, item := range receipt.Items {
		if len(item.Description) > r.MaxLength {
			return r.Points, nil
		}
	}
	return 0, nil
}

// purchaseDayOddRule checks if the purchase day is odd
type purchaseDayOddRule struct {
	Points int `json:"points"`
}

func (r *purchaseDayOddRule) Evaluate(receipt model.Receipt) (int, error) {
	if receipt.PurchaseTime.Day()%2 != 0 {
		return r.Points, nil
	}
	return 0, nil
}

// purchaseTimeRangeRule checks if the purchase time is within a specific range
type purchaseTimeRangeRule struct {
	StartHour int `json:"startHour"`
	EndHour   int `json:"endHour"`
	Points    int `json:"points"`
}

func (r *purchaseTimeRangeRule) Evaluate(receipt model.Receipt) (int, error) {
	if receipt.PurchaseTime.Hour() >= r.StartHour && receipt.PurchaseTime.Hour() <= r.EndHour {
		return r.Points, nil
	}
	return 0, nil
}