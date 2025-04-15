package processor

import (
	"encoding/json"
	"fmt"
	"github.com/vijaya-chintapalli/Projects/model"
)

// ruleProcessor is responsible for storing and processing rules
type ruleProcessor struct {
	storeNameRules   []StoreNameRule
	itemMatchRules   []ItemMatchRule
}

// NewProcessor initializes a new rule processor
func NewProcessor() *ruleProcessor {
	return &ruleProcessor{}
}

// AddRule adds a rule to the processor
func (rp *ruleProcessor) AddRule(ruleType model.RuleType, ruleDefinition string) error {
	var err error

	// Use the constants directly here
	switch ruleType {
	case model.RuleTypeStoreName:
		var storeNameRule StoreNameRule
		err = json.Unmarshal([]byte(ruleDefinition), &storeNameRule)
		if err == nil {
			rp.storeNameRules = append(rp.storeNameRules, storeNameRule)
		}
	case model.RuleTypeItemMatch:
		var itemMatchRule ItemMatchRule
		err = json.Unmarshal([]byte(ruleDefinition), &itemMatchRule)
		if err == nil {
			rp.itemMatchRules = append(rp.itemMatchRules, itemMatchRule)
		}
	default:
		return fmt.Errorf("unsupported rule type: %v", ruleType)
	}

	if err != nil {
		return fmt.Errorf("failed to parse rule definition: %v", err)
	}

	return nil
}

// Process applies all the rules to the receipt and returns the total points
func (rp *ruleProcessor) Process(receipt model.Receipt) (int, error) {
	totalPoints := 0

	// Apply StoreName Rules
	for _, rule := range rp.storeNameRules {
		if receipt.StoreName == rule.Value {
			totalPoints += rule.Points
		}
	}

	// Apply ItemMatch Rules
	for _, rule := range rp.itemMatchRules {
		for _, item := range receipt.Items {
			if contains(rule.IDs, item.ID) {
				totalPoints += int(float64(item.Price) * rule.Rate)
			}
		}
	}

	return totalPoints, nil
}

// StoreNameRule checks if the retailer name matches and applies points
type StoreNameRule struct {
	Value  string `json:"value"`
	Points int    `json:"points"`
}

// ItemMatchRule applies a multiplier for specific items
type ItemMatchRule struct {
	IDs   []string `json:"ids"`
	Rate float64   `json:"rate"`
}

// Helper function to check if an item ID exists in a slice
func contains(ids []string, id string) bool {
	for _, val := range ids {
		if val == id {
			return true
		}
	}
	return false
}