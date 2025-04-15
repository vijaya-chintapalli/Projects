package main

import (
    "receipt-rule-engine-challenge/model" // Ensure that model is properly imported
    "fmt"
)

func main() {
    // Example of creating an Item and a Receipt
    item := model.Item{
        ID:    "123",
        Price: 10.50,
    }
    fmt.Println(item)
}