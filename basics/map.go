package main

import "fmt"

func main()  {
	// Using make function to create map
	scores:=make(map[string]int)
	scores["Alice"]=23
	scores["Bob"]=25
	scores["Carol"]=27
	for name, score := range scores {
		println(name, score)
	}
	// Using the map literal syntax 
	scores2:= map[string]int{
		"Alice": 23,
		"Bob": 25,
		"Carol": 20,
	}
	for name, score := range scores2 {
		fmt.Println(name, score)
	}
	//Accessing map elements
	fmt.Println(scores["Alice"]) // Outputs: 23
	fmt.Println(scores["Bob"]) // Outputs: 25

	// Accessing a map element with a key that does not exist
	value, exists := scores["Peter"]
	if exists {
		fmt.Println("Peter's score:", value)
	} else {
		fmt.Println("Peter's score does not exist")
	}
	// Deleting a map element
	delete(scores, "Alice")
	
}
