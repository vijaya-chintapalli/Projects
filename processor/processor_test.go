package processor

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vijaya-chintapalli/Projects/model"
)

func TestProcessor(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		p := NewProcessor()

		err := p.AddRule(model.RuleTypeStoreName, `{"value":"Target","points":100}`)
		require.NoError(t, err)
		err = p.AddRule(model.RuleTypeItemMatch, `{"ids":["111"],"rate":0.1}`)
		require.NoError(t, err)

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

		// Print the points and receipt details
		points, err := p.Process(receipt)
		require.NoError(t, err)
		fmt.Printf("Processed Points: %d\n", points)
		require.Equal(t, 1325, points)
	})
}
