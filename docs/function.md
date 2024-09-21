---
layout: page
title: Functions
permalink: /functions/
nav_order: 7

---

## Functions
A function is a block of code that can be reused. Along with being modular and reusable, functions also benefit the readability of code. That is, rather than a collection of loosely-connected lines of code, bundling your code into a function helps readers of your code determine its intended logic. It helps reduce complexity by decomposing it into simpler, related pieces. The first function you were exposed to is the main function which was introduced under the program structure section of this document. To define a function, you use the func keyword followed by the function name, the parameter list, return types and the body of the function.
The signature of a function is as follows 
```go
func name(parameter-list) (result-list) {
    body
}
```
 For example, the following code defines a function named add that takes two integers as parameters and returns their sum:

 ```go
 func add(x int, y int) int {
    return x + y
}

 ```

Here are the various parts of a Go function:
- ***Function name***: The function name is a unique identifier that is used to call the function.
- ***Parameter list***: The parameter list is a list of names and types of the variables that are passed to the function. This values are generally supplied by the caller
- ***Return list***: The list of types that are returned from the function
- ***Body***: The body of the function is the block of code that is executed when the function is called.


The function name must be unique within the scope in which it is declared. The parameter list is optional. If there are no parameters, the parameter list is empty. The return list is optional as well if the function doesn't return anything is is left out. The body of the function is required.
Above is an example of a function with name add and parameter list  of x and y which are of type int. return type of int. The function body is everything inside of the curly braces in our example above it is the line return x + y

To invoke, call a function, you use the function name followed by the arguments. For example, the following code calls the add function with the arguments 10 and 20:
```go
z := add(10, 20)
```
Here is another function with no paramater list and no return list 
```go
func sayHello() {
    fmt.Println("Hello, world!")
}

```


Functions can be used to:

1. Break down large tasks into smaller, more manageable pieces.
2. Avoid duplicating code.
3. Improve readability and maintainability of code.

Here are some of the key features of Golang functions:
1. Functions are first-class citizens: This means that functions can be assigned to variables, passed as arguments to other functions, and returned from other functions.
2. Functions are closures: This means that a function can access the variables of the enclosing scope, even if those variables are declared after the function is defined.
3. Functions are concurrent: This means that multiple functions can be executed at the same time.


```go
package main

import "fmt"

func main() {
	// call add function with argument of 10 and 20
	z := add(10, 20)
	// print th evalue of variable z which is the result of add function
	fmt.Println(z)
}

// implementation of function add which takes two arguments and return their sum
func add(x int, y int) int {
	return x + y
}

// this will print 30
```

Arguments are passed by value: the function receives a copy of each argument; modifications to the copy do not affect the caller. However, if the argument contains some kind of reference( see reference type in the datatype section), like a pointer, slice, map, function, or channel, then the caller may be affected by any modifications the function makes to variables indirectly referred to by the argument.

You may have noticed that the main function in Go programs doesn't accept any arguments, nor does it actually return anything either! Ultimately, it's up to you, the programmer, to decide what you need from the functions you write it. As a rule of thumb, think about what inputs (if any) are needed for the function to perform its tasks and what output result (again, if any) you want to get as the result of invoking your function. Generally, functions are handy and help you reuse code for different purposes, e.g., a different set of arguments can be supplied
