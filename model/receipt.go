package model

import "time"

// Receipt represents a receipt with items, total amount, etc.
type Receipt struct {
    ID           string    `json:"id"`
    StoreName    string    `json:"storeName"`
    PurchaseTime time.Time `json:"purchaseTime"`
    Items        []Item    `json:"items"`
    Total        float64   `json:"total"`
}
