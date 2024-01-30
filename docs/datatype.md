---
layout: page
title: Data type
permalink: /data-type/
nav_order: 6

---


## Data type

Golang has a variety of data types.
Go’s types fall into four categories:

1. Basic types: numbers, strings, and booleans.
2. Aggregate types: arrays and structs.
3. Reference types: pointers, slices, maps, functions, and channels .
4. Interface types

 The most common Basic data types are:
- Integers: Integers can store whole numbers.
- Floating-point numbers: Floating-point numbers can store numbers with decimal points.
- Strings: Strings can store sequences of characters.
- Booleans: Booleans can store values of true or false.

Here are some of the key features of Golang data types:

- Type inference: Golang supports type inference, which means that you do not need to specify the type of a variable when you declare it.
- Type safety: Golang is a type-safe language, which means that the compiler will check that the types of values are compatible before executing your code.
- Type conversion: Golang supports type conversion, which allows you to convert values from one type to another.
- Data types are strongly typed: This means that the type of a variable cannot be changed once it has been declared.
- Data types are immutable: This means that the value of a variable cannot be changed once it has been assigned.
- Data types are scoped: This means that the names of data types are only visible within the scope in which they are declared.
- Composite types: Golang supports composite types, which are made up of other data types. For example, a struct is a composite type that can store multiple data types.
- Pointers: Golang supports pointers, which are variables that store the address of another variable.
- Interfaces: Golang supports interfaces, which are a way of defining a contract that a type must fulfill.

```go
package main

import "fmt"

func main() {
	// declare and integer variable x
	var x int
	// assign the value 10 to the variable x
	x = 10
	// print the value of the variable x
	fmt.Println(x)
	// declare a floating-point variable named y
	var y float64
	// assign the value 10.5 to the variable y
	y = 10.5
	//print the value of variable y
	fmt.Println(y)
	// declare a string variable named z
	var z string
	// assign the value Hello, World! to z
	z = "Hello, World!"
	//print the value of variable z
	fmt.Println(z)
	// declare a boolean variable named a
	var a bool
	// assign the value true to the variable a
	a = true
	// print the value of a
	fmt.Println(a)
}
/* the above outputs
10
10.5
Hello, World!
true
*/
```
