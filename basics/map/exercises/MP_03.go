package main

import (
	"fmt"
	"strings"
)

// Write a program that counts the frequency of each word in a given text and stores the results in a map.
func countWordFrequency(text string) map[string]int {
 frequency := make(map[string]int)

 // Split the text into words
 words := strings.Fields(text)

 // Iterate over each word
 for _, word := range words {
  // Update the frequency count
  frequency[word]++
 }

 return frequency
}

func main() {
 // Example usage
 text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
 wordFrequency := countWordFrequency(text)

 // Print the word frequencies
 for word, freq := range wordFrequency {
  fmt.Printf("%s: %d\n", word, freq)
 }
}