package main

import (
    "fmt"
    "github.com/vijaya-chintapalli/Projects/model"  // Correctly importing the model package
)

func main() {
    item := model.Item{
        Name:        "Example Item",
        Description: "An example description",
    }
    fmt.Println(item)
}
