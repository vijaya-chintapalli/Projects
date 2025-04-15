package model

// Rule interface represents a rule that can be evaluated on a receipt
type Rule interface {
    Evaluate(receipt Receipt) (int, error)
}
