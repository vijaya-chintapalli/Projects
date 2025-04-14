package processor

import (
	"receipt-rule-engine-challenge/model"
)

// RuleProcessor A Rule-Based receipt processor
type RuleProcessor interface {

	// AddRule adds an arbitrary rule to the processor to be used in handling all subsequent receipts.
	//
	// ruleType refers to the RuleType constant classifying the rule.
	// ruleDefinition is expected to be a JSON string that can be parsed into a rule of the indicated type.
	AddRule(ruleType model.RuleType, ruleDefinition string) error

	// Process evaluates a receipt against the available rules, and returns the number of points to award.
	Process(receipt model.Receipt) (int, error)
}

func NewProcessor() RuleProcessor {
	// TODO: Initialize processor
	return nil
}
