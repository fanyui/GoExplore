package main

import "fmt"
func main()  {
	fmt.Println("Hello World")
	var name string = "John"
	var age int = 20
	var isCool = true
	// Shorthand
	// name := "John"
	size := 1.3
	name, email := "John", "john@gmail.com"
	fmt.Printf("Name is %v\n age is %v\n email is %v\n",name, age, email)
	fmt.Printf("isCool is %v\n size is %v\n",  isCool, size)

}