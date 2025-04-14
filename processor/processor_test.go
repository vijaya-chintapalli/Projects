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
		p := NewProcessor()

		require.NoError(t, p.AddRule(model.RuleTypeStoreName, `{"value":"Target","points":100}`))
		require.NoError(t, p.AddRule(model.RuleTypeItemMatch, `{"ids":["111"],"rate":0.1}`))

		receipt := model.Receipt{
			ID:           uuid.NewString(),
			StoreName:    "Target",
			PurchaseTime: time.Now(),
			Items: []model.Item{
				{ID: "111", Price: 12.25},
				{ID: "222", Price: 20.99},
			},
			Total: 32.24,
		}

		points, err := p.Process(receipt)
		require.NoError(t, err)

		// 100 Points for a Target receipt
		// 1225 Points for is 10% of item '111'
		expectedPoints := 100 + 1225

		require.Equal(t, expectedPoints, points)
	})
}
