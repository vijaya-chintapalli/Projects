package model

import (
"time"
"encoding/json"
"fmt")

// Receipt is a simplified receipt for the purposes of this challenge. It represents a paper receipt snapped by one
// of our users and sent in for points. The attributes of this receipt represent the values read from the receipt
// except where otherwise noted.
type Receipt struct {

	// ID represents a unique identifier for the receipt. This is an internal Fetch ID for tracking this receipt.
	ID string

	StoreName    string
	PurchaseTime time.Time
	Total        float64

	Items []Item
}

type Item struct {
	// ID An identifier created by Fetch. Similar to a barcode, it uniquely and accurately identifies the item according
	// to our internal product catalog.
	ID string

	Price float64
}
func ParseReceipt(data []byte) (*Receipt, error) {
	var receipt Receipt
	if err := json.Unmarshal(data, &receipt); err != nil {
		return nil, fmt.Errorf("failed to parse receipt: %w", err)
	}
	return &receipt, nil
}
func ParsePurchaseTime(timeStr string) (time.Time, error) 
	return time.Parse("2006-01-02 15:04:05", timeStr)
}

func SerializeReceipt(receipt *Receipt) ([]byte, error) {
	return json.Marshal(receipt)
}

func SerializeItem(item *Item) ([]byte, error) {
	return json.Marshal(item)
}
