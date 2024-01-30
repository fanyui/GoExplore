---
layout: page
title: Variable
permalink: /variable/
nav_order: 5

---


## Variable
 They are named memory locations that can hold values. Variables are declared using the var keyword, followed by the variable name and the data type. Variables are used to store data in Golang. They can be used to store any type of data, including numbers, strings, and objects, etc. Variables are a fundamental part of Golang programming, and they are used in almost every piece of code. We have already made use of variables in most of the sections above but we will dive deeper into it here. 

>***Note***: Identifiers and variables are related concepts but serve different purposes in the context of programming. with difference being that
An identifier is a name given to various program entities, including variables, functions, packages, constants, and types. while A variable is a storage location in the computer's memory that is identified by a name (identifier) and used to store and manipulate data during program execution.

Just like we explained in the identifire section,
- Variable names must start with a letter or an underscore.
- Variable names can contain letters, numbers, and underscores.
- Variable names cannot start with a number.
- Variable names cannot contain spaces.
- Variable names cannot be the same as keywords.

### How to declare variables

Variables are declared using the var keyword, followed by the variable name and the data type. Variable declaration create one or more variables, ties it to the identifier and then gives it a type and an initial value.For example, the following code declares a variable named x of type int:
```go
var x int        // variable x of type int
var A, B float64 // variable A, B of type float
var c = 7        // variable c inferred type int
var d string     // variable d of type string
var e, f, g = true, 2.3, "four" // variables e of type boolean, f of type float64, g  of type string

i := 100         // variable i of type inferred int
```

Using the examples above you will notice that the long and full annotated version of a variable declaration is typically of the form 
`var name type = expression`
however there exist short hand notations eg the third example above `var c = 7 ` 7 is inferred and the last line `i := 100`  declars and assigned 100 to i. In this case the type of i is also inferred from the usage.
Variable declaration specifies the type a variable can store. 
Integers can store whole numbers. Floating-point numbers can store numbers with decimal points. Strings can store sequences of characters. Booleans can store values of true or false.
Once a variable is declared, it therefore means that we can give it a value in otherword we can asign a value to it. We can only assign values of the particular type to the value. hence you cannot assign a floating point number to a variable that was declared as int. 
To asign a value to a variable, we use the equal(`=`) operator To do this, For example, the following code assigns the value 10 to the variable x:
`x = 10` since we declared x above to be of type int and 1 being an integer means we can perform that assignment.

### Printing variables
As with most printing, the fmt package allows us to display the content of a variable. The following code outputs the value stored in variable x.
```go
fmt.Println(x)
```

### using vairables in expression
Variables can be used in expressions. For example, the following code evaluates the expression x + 1 and assigns the result to the variable c:
```
c = x + 1
```

Variables have a scope which determines where they can be used. The scope of a variable is determined by where it is declared. Variables declared at the package level have the widest scope. Variables declared inside a function have a narrower scope. Variables declared inside a block have the narrowest scope

```go
package main

import "fmt"

func main() {
	var x int
	var A, B float64
	var c = 7
	var d string
	var e, f, g = true, 2.3, "four"
	x = 10
	A = 2.5
	B = 3.5
	d = "hello"
	fmt.Printf("x = %d, A = %f,B = %f, d = %s, e = %v, f = %f, g = %s, c = %d \n", x, A, B, d, e, f, g, c)
}

// the above prints: x = 10, A = 2.500000,B = 3.500000, d = hello, e = true, f = 2.300000, g = four, c = 7
```
