package model

// Item represents an individual item in the receipt.
type Item struct {
    ID          string  `json:"id"`          // Item ID
    Description string  `json:"description"` // Item description (if needed for length checks or other logic)
    Price       float64 `json:"price"`       // Item price
}