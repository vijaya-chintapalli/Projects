package model

import (
    "time"
    "encoding/json"
    "fmt"
)

// Receipt represents a simplified receipt for the purposes of this challenge.
type Receipt struct {
    ID           string
    StoreName    string
    PurchaseTime time.Time
    Total        float64
    Items        []Item // Reference to Item struct from item.go
}

func ParseReceipt(data []byte) (*Receipt, error) {
    var receipt Receipt
    if err := json.Unmarshal(data, &receipt); err != nil {
        return nil, fmt.Errorf("failed to parse receipt: %w", err)
    }
    return &receipt, nil
}

func ParsePurchaseTime(timeStr string) (time.Time, error) {
    return time.Parse("2006-01-02 15:04:05", timeStr)
}

func SerializeReceipt(receipt *Receipt) ([]byte, error) {
    return json.Marshal(receipt)
}

func SerializeItem(item *Item) ([]byte, error) {
    return json.Marshal(item)
}