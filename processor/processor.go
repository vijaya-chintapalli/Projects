package processor

import (
	"receipt-rule-engine-challenge/model"
	"encoding/json"
	"errors"
	"fmt"
	"receipt-rule-engine-challenge/model"
	"strings"
	"time"
)

type ruleProcessor struct {
	rules map[model.RuleType]Rule
}
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
	return &ruleProcessor{
		rules: make(map[model.RuleType]Rule),
	}
}

func (rp *ruleProcessor) AddRule(ruleType model.RuleType, ruleDefinition string) error {
	var rule Rule
	var err error

	switch ruleType {
	case model.RetailerNameAlpha:
		rule = &retailerNameAlphaRule{}
	case model.TotalRoundDollar:
		rule = &totalRoundDollarRule{}
	case model.TotalMultipleOfQuarter:
		rule = &totalMultipleOfQuarterRule{}
	case model.ItemCountMultiplier:
		rule = &itemCountMultiplierRule{}
	case model.ItemDescriptionLength:
		rule = &itemDescriptionLengthRule{}
	case model.PurchaseDayOdd:
		rule = &purchaseDayOddRule{}
	case model.PurchaseTimeRange:
		rule = &purchaseTimeRangeRule{}
	default:
		return fmt.Errorf("unsupported rule type: %v", ruleType)
	}

	if err := json.Unmarshal([]byte(ruleDefinition), rule); err != nil {
		return fmt.Errorf("failed to unmarshal rule definition: %w", err)
	}

	rp.rules[ruleType] = rule
	return nil
}

func (rp *ruleProcessor) Process(receipt model.Receipt) (int, error) {
	totalPoints := 0

	for ruleType, rule := range rp.rules {
		points, err := rule.Evaluate(receipt)
		if err != nil {
			return 0, fmt.Errorf("error evaluating rule %v: %w", ruleType, err)
		}
		totalPoints += points
	}

	return totalPoints, nil
}

// Rule Implementations

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

type totalRoundDollarRule struct {
	Points int `json:"points"`
}

func (r *totalRoundDollarRule) Evaluate(receipt model.Receipt) (int, error) {
	if receipt.Total == float64(int64(receipt.Total)) {
		return r.Points, nil
	}
	return 0, nil
}

type totalMultipleOfQuarterRule struct {
	Points int `json:"points"`
}

func (r *totalMultipleOfQuarterRule) Evaluate(receipt model.Receipt) (int, error) {
	if receipt.Total == 0 {
		return 0, nil
	}
	if remainder := receipt.Total - float64(int64(receipt.Total/0.25))*0.25; remainder < 0.001 {
		return r.Points, nil
	}
	return 0, nil
}

type itemCountMultiplierRule struct {
	Multiplier int `json:"multiplier"`
}

func (r *itemCountMultiplierRule) Evaluate(receipt model.Receipt) (int, error) {
	return len(receipt.Items) / 2 * r.Multiplier, nil
}

type itemDescriptionLengthRule struct {
	RequiredLength int `json:"requiredLength"`
	Points         int `json:"points"`
}

func (r *itemDescriptionLengthRule) Evaluate(receipt model.Receipt) (int, error) {
	points := 0
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ID)
		if len(trimmed)%r.RequiredLength == 0 {
			points += r.Points
		}
	}
	return points, nil
}

type purchaseDayOddRule struct {
	Points int `json:"points"`
}

func (r *purchaseDayOddRule) Evaluate(receipt model.Receipt) (int, error) {
	if receipt.PurchaseTime.Day()%2 != 0 {
		return r.Points, nil
	}
	return 0, nil
}

type purchaseTimeRangeRule struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	Points   int    `json:"points"`
}

func (r *purchaseTimeRangeRule) Evaluate(receipt model.Receipt) (int, error) {
	purchaseTime := receipt.PurchaseTime
	startTime, err := time.Parse("15:04", r.Start)
	if err != nil {
		return 0, fmt.Errorf("invalid start time format: %w", err)
	}

	endTime, err := time.Parse("15:04", r.End)
	if err != nil {
		return 0, fmt.Errorf("invalid end time format: %w", err)
	}

	// Compare just the time components
	pt := time.Date(0, 0, 0, purchaseTime.Hour(), purchaseTime.Minute(), 0, 0, time.UTC)
	st := time.Date(0, 0, 0, startTime.Hour(), startTime.Minute(), 0, 0, time.UTC)
	et := time.Date(0, 0, 0, endTime.Hour(), endTime.Minute(), 0, 0, time.UTC)

	if pt.After(st) && pt.Before(et) {
		return r.Points, nil
	}
	return 0, nil
}
