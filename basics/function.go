package main

import "fmt"
func main() {
    // Call the `add` function with the arguments `10` and `20`.
    z := add(10, 20)
    // Print the value of the variable `z`.
    fmt.Println(z)
}

// Declare a function named `add` that takes two integers as parameters and returns their sum.
func add(x int, y int) int {
	return x + y
}