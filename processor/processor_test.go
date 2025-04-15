package processor

import (
	"github.com/vijaya-chintapalli/Projects/model"
	"testing"
	"github.com/stretchr/testify/require"
	"time"
)

func TestProcessor(t *testing.T) {
	p := NewProcessor()

	// Adding rules
	err := p.AddRule(model.RuleTypeStoreName, `{"value":"Target","points":100}`)
	require.NoError(t, err)

	err = p.AddRule(model.RuleTypeItemMatch, `{"ids":["111"],"rate":0.1}`)
	require.NoError(t, err)

	// Creating a sample receipt
	receipt := model.Receipt{
		ID:           "1234",
		StoreName:    "Target",
		PurchaseTime: time.Now(),
		Items: []model.Item{
			{ID: "111", Price: 12.25, Description: "Item A"},
			{ID: "222", Price: 20.99, Description: "Item B"},
		},
		Total: 32.24,
	}

	// Processing the receipt
	points, err := p.Process(receipt)
	require.NoError(t, err)

	// Expected points: 100 points from the store match, and 1.225 points from the item match rule (12.25 * 0.1)
	require.Equal(t, 101, points)
}
