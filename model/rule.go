package model

// Rule defines a generic interface that any rule should implement
type Rule interface {
	Evaluate(receipt Receipt) (int, error)
}