package main

import "fmt"

//Write a function in Go that takes a map of student names and their scores as input and returns the average score.
func calculateAverageScore(scores map[string]int) float64 {
    sum := 0
    count := 0

    for _, score := range scores {
        sum += score
        count++
    }

    if count == 0 {
        return 0.0
    }

    average := float64(sum) / float64(count)
    return average
}

func main() {
    // Example usage
    studentScores := map[string]int{
        "Alice": 90,
        "Bob":   85,
        "Charlie": 95,
        "Dave":  78,
    }

    averageScore := calculateAverageScore(studentScores)
    fmt.Println("Average Score:", averageScore)
}
