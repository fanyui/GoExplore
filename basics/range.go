package main

import "fmt"

func main() {
    // Create an array of integers
    var ages = [5]int{18, 21, 25, 30, 35}
    // Iterate over the elements of the array
    for i, age := range ages {
        fmt.Println(i, age)
    }
    // Create a map of strings to integers
    var names = map[string]int{"Alice": 25, "Bob": 30, "Carol": 35}
    // Iterate over the keys of the map
    for name, age := range names {
        fmt.Println(name, age)
    }
}