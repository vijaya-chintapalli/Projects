package processor

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"receipt-rule-engine-challenge/model"
	"testing"
	"time"
)

func TestProcessor(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		// Create a new processor instance
		p := NewProcessor()

		// Add rules: Store Name and Item Match
		require.NoError(t, p.AddRule(model.RuleTypeStoreName, `{"value":"Target","points":100}`))
		require.NoError(t, p.AddRule(model.RuleTypeItemMatch, `{"ids":["111"],"rate":0.1}`))

		// Create a sample receipt
		receipt := model.Receipt{
			ID:           uuid.NewString(),
			StoreName:    "Target",
			PurchaseTime: time.Now(),
			Items: []model.Item{
				{ID: "111", Price: 12.25}, // This item should match RuleTypeItemMatch
				{ID: "222", Price: 20.99}, // This item should not match any rule
			},
			Total: 32.24,
		}

		// Process the receipt and calculate the points
		points, err := p.Process(receipt)
		require.NoError(t, err)

		// Points calculation:
		// 100 Points for Target store
		// 10% of 12.25 (price of item "111") = 1.225, which equals 1225 points (multiplied by 100)
		expectedPoints := 100 + 1225 // 100 points for Target + 1225 points for item "111"

		// Assert that the calculated points match the expected points
		require.Equal(t, expectedPoints, points)
	})
}