package main

import (
	"fmt"
	"strconv"
)
type Person struct {
	Name   string
	Age    int
	Height float64
}
func (p Person) String() string {
	return  p.Name + " is " + strconv.Itoa(int(p.Age)) + " years old and " + strconv.FormatFloat(p.Height,'E', -1,32) + " feet tall"
}
func main()  {

	person := Person{
    Name:   "John Doe",
    Age:    30,
    Height: 5.8,
	}
	// Accessing struct fields
	fmt.Println(person.Name)    // Outputs: John Doe
	fmt.Println(person.Age)     // Outputs: 30
	fmt.Println(person.Height)  // Outputs: 5.8

	// Updating struct fields
	person.Name = "Harisu Fanyui"
	person.Age = 28
	fmt.Println("Updated Name:", person.Name) // Outputs: Updated Name: Harisu Fanyui
	fmt.Println("Updated Age:", person.Age)   // Outputs: Updated Age: 28
	fmt.Println(person.String())
}