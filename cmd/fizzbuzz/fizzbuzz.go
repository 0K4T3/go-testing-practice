package main

import (
    "fmt"

    "go-testing-practice/pkg/fizzbuzz"
)

func main() {
    for i := 1; i < 101; i++ {
        fmt.Println(fizzbuzz.FizzBuzz(i))
    }
}
