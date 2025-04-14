package model

import "time"

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
