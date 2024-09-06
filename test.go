package main

import (
    "fmt"
    "os"
)

func main() {
    testVar := os.Getenv("TEST_VAR")
    if testVar == "" {
        fmt.Println("TEST_VAR is not set")
    } else {
        fmt.Printf("TEST_VAR=%s\n", testVar)
    }
}
