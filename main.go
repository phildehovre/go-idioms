package main

import (
    "fmt"
)

func main() {
    // Initializing a map with make
    m := make(map[string]int)

    // Checking if a key exists in the map
    if _, ok := m["key"]; !ok {
        fmt.Println("Key does not exist")
    }

    // Iterating over a map
    for key, value := range m {
        fmt.Println(key, value)
    }

    // Deferring a function call
    defer fmt.Println("Deferred function call")

    // Handling multiple return values
    result, err := someFunction()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}

func someFunction() (string, error) {
    return "Result", nil
}