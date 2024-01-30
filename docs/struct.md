---
layout: page
title: Struct
permalink: /struct/
nav_order: 11

---



## Structs

Analogically speaking, if you have programmed in other Object-oriented supported languages, you may be familiar with classes. the general definition of a class will be that it's a blueprint or template from which objects can be created. Like classes in Java or other Object Oriented programming languages, Struct is a composite data type blueprint that allows you to group different types into a single unit. They are very similar to classes but they don't have the methods as classes will have. Struct provides a way of creating complex data structures from custom types

### How to Create a Struct  
To create a struct, we use the type keyword followed by the name of the struct and then, subsequently, the fields it will contain. 
```go
type Person struct {
    Name string
    Age int
    Height float64
}
```
Now we need to instantiate an object of the type Person we defined above.
```go
person := Person{
    Name:   "John Doe",
    Age:    30,
    Height: 5.8,
}
```
with the above, we can now access or modify the fields on the person above.
```go
package main

import "fmt"

func main() {
	type Person struct {
		Name   string
		Age    int
		Height float64
	}
	person := Person{
		Name:   "John Doe",
		Age:    30,
		Height: 5.8,
	}
	//Accessing struct fields
	fmt.Println(person.Name) 	// Outputs: John Doe
	fmt.Println(person.Age)  	// Output: 30
	fmt.Println(person.Height)	// Output: 5.8

	// Updateing struct fields
	person.Name = "Harisu Fanyui"
	person.Age = 28
	fmt.Println("Updated Name: ", person.Name) // Outputs: Updated Name: Harisu Fanyui
	fmt.Println("Updated Age: ", person.Age)   // Outputs: Updated Age: 28

}
```

It is important to note that structs can also have methods. These can come in two forms. 
1. Pointer Receiver 
2. Value receiver. 

Consider using a value receiver only when no mutation is happening on the struct. ie say just some computation needed to generate some result of the struct. eg let's write a function that returns the string version of an object. Just like we have the toString method in other object-oriented programming languages. here we can have it written as follows
```go
func (p Person) String() string {
	return p.Name + " is " + strconv.Itoa(int(p.Age)) + " years old and " + strconv.FormatFloat(p.Height, 'E', -1, 32) + " feet tall"
}

```

however, for the pointer receiver, you want to use it for two reasons. 
- When there is a mutation happening in the method on the receiver object. 
- To avoid copying data on each method call.

####  Excercise

1. Write a function that takes a struct representing a Rectangle (with fields for length and width) as input and calculates its area. 
2. Create a struct to represent a Book, with fields for title, author, and price. Write a function that takes a slice of Book and returns the total price of all the books. 

More in Miscellanous
