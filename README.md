# GoExplore

# Introduction to Golang

## Brief Overview of Go and its History.

Go is a programming language that was created in 2009 by Google engineers Robert Griesemer, Rob Pike, and Ken Thompson. It was designed to be a fast, efficient, and easy-to-use language for building concurrent and networked software.

Go was inspired by several other programming languages, including C, Pascal, and Oberon. Its syntax is similar to that of C, but it includes several features that make it easier to write safe and efficient code, such as garbage collection, memory safety, and built-in concurrency support.

One of the key features of Go is its simplicity. It has a small and concise syntax makes it easy to learn and read. It also strongly focuses on performance, with a compiler that generates fast and efficient code. It is a statically typed(which means that the types of variables and expressions are known at compile time), compiled language(which means that it is converted to machine code before it is executed) that is designed to be simple, reliable, and efficient. Golang is often used for web development, command-line tools, and data science. It is a popular choice for both beginners and experienced developers.
Today, there are estimated to be over 1.1 million developers who choose Go as their primary language!

## Why should you learn Golang?(Reasons for Learning Go).
There are many reasons to learn Golang. Here are a few of the most important:
- It's a great language for building networked and concurrent software, which is becoming increasingly important in today's world.
- Golang is a simple language to learn. The syntax is straightforward and there are not many keywords.
- Golang is a reliable language. It is designed to be robust and to handle errors gracefully.
- Golang is an efficient language. It is compiled to machine code, which makes it fast and efficient.
- Golang is a versatile language. It can be used for a variety of tasks, including web development, command-line tools, and data science.
- Go is particularly well-suited for building web servers and microservices, as well as for working with big data and machine learning. It's also used in industries such as finance, healthcare, and gaming.


### What are the benefits of using Golang?
Some of the few benefits you will get are
- ***Speed***: Golang programs are compiled to machine code, which makes them fast and efficient.
- ***Reliability***: Golang is designed to be robust and to handle errors gracefully.
- ***Simplicity***: Golang is a simple language to learn and use.
- ***Community***: Golang has a large and active community of developers.
- ***Garbage collection***: Go is also great at memory management, and this is implemented in Go under the hood through its use of a garbage collector

# Basics
## Getting started

Learning a new language for the first time is no easy feat! However, keep in mind that every Go developer, no matter how skilled, was once in your shoes: learning the syntax and inner workings of the language for the first time.


Be sure to make use of all the resources available to you, and maintain a constant rhythm. Throughout this book, you'll have many opportunities to practice the Go skills you'll pick up along the way (especially those that you'll need for your entire go career!). By leveraging your creativity and problem-solving skills, I know you'll find development with Go rewarding. Good luck!

## Step-by-Step Guide for Installing Go

Getting Go installed on your local machine is straightforward. If you're using Linux, Mac, or Windows, you can download and execute an installer to get up and running immediately.

### Installing Go on Windows

To install Go on Windows, follow these steps:
1. Download the Windows installer from the official Go website: https://golang.org/dl/
2. Double-click on the installer to start the installation process.
3. Follow the prompts to complete the installation. Choose the default options unless you have a specific reason to change them.
4. Once the installation is complete, open the Command Prompt or PowerShell and type go version to verify that Go is installed correctly.

### Installing Go on Mac
To install Go on a Mac, follow these steps:
1. Download the macOS installer from the official Go website: https://golang.org/dl/
2. Double-click on the installer to start the installation process.
3. Follow the prompts to complete the installation. Choose the default options unless you have a specific reason to change them.
4. Once the installation is complete, open the Terminal app and type go version to verify that Go is installed correctly.


### Installing Go on Linux
To install Go on Linux, follow these steps:
1. Open a terminal window.
2. Update the package list by typing `sudo apt update` (for Debian/Ubuntu) or `sudo yum update` (for CentOS/Fedora).
3. Install the Go package by typing `sudo apt install golang`(for Debian/Ubuntu) or `sudo yum install golang` (for CentOS/Fedora).
4. Once the installation is complete, type `go version` to verify that Go is installed correctly.

If you still have issues after following the instructions above, you can checkout the official website [Download and Install](https://go.dev/doc/install)

## Setting Up Environment Variables
After the installation steps covered above,it's time for us to set up environment variables necessary for running go programs and using go tools.

### Setting Up Environment Variables on Windows
1. Open the Control Panel and go to System and Security > System > Advanced system settings.
2. Click on the Environment Variables button.
3. Under System Variables, click on New.
4. Enter the following information:
   - Variable name: `GOROOT`
   - Variable value: `C:\Go` (or the directory where Go is installed)
5. Click on OK to save the variable.
6. Under System Variables, click on New again.
7. Enter the following information:
   - Variable name: `GOPATH`
   - Variable value:`%USERPROFILE%\go` (or the directory where you want to store your Go projects)
8. Click on OK to save the variable

### Setting Up Environment Variables on Mac/Linux
Here the steps are both same for Mac and linux

1. Open the Terminal app.
2. Type `nano ~/.bash_profile` to open the Bash profile file.
3. Add the following lines to the file:
   `export GOROOT=/usr/local/go`
   `export GOPATH=$HOME/go`
   `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`
4. Press `Ctrl+X`, then `Y`, then Enter to save the file and exit Nano.
5. Type source `~/.bash_profile` to reload the Bash profile.


## Program Structure. 
Without wasting much time, We will start straight away by an example Hello World program in go and then explain what each part of the program represent. Then we move forward to defining some terms which apply to this and will apply to all other Go programs you write. If you this is your first programming language then you can simply flow along however, if you are coming from a different programming language, you will have to be more flexible to adapt to the new ways in which go programs are written. This is because it diverts from the way other progamming languages do operate. 

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
Follow the instructions below to get started with your first  go programe 
1. Create a folder called hello `mkdir hello`
2. Inside the folder create a file called helloworld.go  `touch helloworld.go`
3. Open the file with name `helloworld.go` and **type** the code above inside and save it.
4. Open a terminal inside the hello folder and type `go run helloworld.go`
5. `Hello, World!` should be displayed on the terminal next line



Like you have read in the overview and history section, Go is a compiled language as such go programs are compiled before run. the go command is available on your computer the moment we installed go on the installation section. 

In the programe above, 
- We Declare a main package `package main` first line. Go code is organized into packages (similar to libraries or modules in other languages). A package is a way to group functions, and it's made up of all the files in the same directory.
- Import the fmt package `import "fmt"`, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go. the import command tells the compiler which package is needed by the program and specifically this file. Rule of thumb is to alway import only the packages you need to use 
- Followed is the main function declaration. `func main()` more about function is reserved for further sections of this document. However keep in mind that the main function is the entry point to every go programe as such that is where our example helloworld programe will start executing. 
- Next is the proper implementation of the message printing `fmt.Println("Hello, World!")`. we use the Println() function from the fmt package we had imported earlier  As explained in the second point above this is a function built into the standard library which comes with the installation of go. We will discuss more about packages later.
- One thing you should also note here is that there is no semi-colon or any special way of ending lines in go. you simply need to move to the next line and continue

This marks the end of our first program in golan. Excited?. There's more to come. 

## Syntax
Golang syntax is designed to be simple and easy to read. The language has a small number of keywords and a regular grammar. This makes it easy to learn Golang syntax, even if you are new to programming.
Golang syntax is based on the C programming language. However, Golang has a number of features that make it more modern and expressive. For example, Golang supports type inference, which means that you do not need to specify the type of a variable when you declare it. Golang also supports multiple return values, which can be used to simplify code. The basic syntax elements are:

- Keywords: Keywords are reserved words that have special meaning in Golang.
- Identifiers: Identifiers are names that are used to identify variables, functions, and other entities.
- Literals: Literals are constants that have a specific value.
- Operators: Operators are used to perform operations on values.
- Statements: Statements are instructions that tell the Golang compiler what to do.
- Comments: Comments are used to make your code more readable. Comments can be single-line or multi-line.

### Keywords
Go has a total of 25 keywords of which are listed below.

| break | default| func | interface | select |
| ------| ------ | ---- |-----------| ------ |
| else   | case   | defer| go        |  map   |
| struct| chan   | const | package  | switch |
| if    | range  | goto| faillthrough| type   |
| var   | for    | continue| return| import  |

### Identifiers
Identifiers are used to identify variables, functions, and other entities in Golang. Identifiers can be made up of letters, numbers, and underscores. Identifiers must start with a letter or underscore

# Literals
Literals are constants that have a specific value. Golang has literals for the following types:
```go
package main

import "fmt"

// Constants identifier MaxRetries
const MaxRetries = 3

// Struct Data Type GithubUser
type GithubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

// Package-Private Variable identifier _internalVariable
var _internalVariable int

// Exported Function idenfier CalculateSum
func CalculateSum(a, b int) int {
	return a + b
}

func main() {
	// Local Variable identifier count
	count := 10

	// printing Using the identifers
	fmt.Println("Count:", count)
	fmt.Println("Max Retries:", MaxRetries)

	// Using the GithubUser Struct Type as for idenfier person
	person := GithubUser{Login: "John", Name: "Doe", PublicRepos: 7}
	fmt.Println("Person:", person)
}

```

In the example program above, the idenfiers are `MaxRetries`, `GithubUser`, `_internalVariable` `CalculateSum`, `count`, Notice that these all follow the rules for identifiers which are 
> - Identifiers must start with a letter (a-z, A-Z) or an underscore _.
> - After the initial character, an identifier may contain letters, digits, or underscores.
> - Identifiers are case-sensitive.
> - Identifiers cannot be the same as Go keywords.

### Literals
Literals are constants that have a specific value. Golang has literals for the following types:
- Integers: Integer literals can be written in decimal, octal, or hexadecimal format.
- Floating-point numbers: Floating-point literals can be written in decimal or scientific format.
- Strings: String literals are enclosed in double quotes.
- Booleans: Boolean literals can be true or false.
- Runes: Rune literals are enclosed in single quotes.

```go
var integerLiteral = 42       // Integer literal
var floatLiteral = 3.14       // Floating-point literal
var hexLiteral = 0x1F         // Hexadecimal literal (31 in decimal)
var octalLiteral = 0755       // Octal literal (493 in decimal)
var stringLiteral = "Hello, World!"  // String literal
var boolLiteralTrue = true    // Boolean literal
var boolLiteralFalse = false   // Boolean literal

ascii := 'a'     // Runes Literal
unicode := 'B'   // Runes Literal
newline := '\n'  // Runes Literal
// when printed using the fmt package  we have the below
fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
fmt.Printf("%d %[1]c %[1]q\n", unicode) // "68 D 'D'"
fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"


```


### Operators


Operators are used to perform operations on values. Golang has operators for the following operations:
- Arithmetic operations: Addition, subtraction, multiplication, division, and modulus.
```
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers
```

- Comparison operations: Equal to, not equal to, less than, less than or equal to, greater than, and greater than or equal to.
```go
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

- Logical operations: And, or, and not.

```go
&&     AND    p && q  is  "if p then q else false"
||     OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
```
- Assignment operations: Assignment, addition assignment, subtraction assignment, multiplication assignment, division assignment, and modulus assignment.
```
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const n float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
person := GithubUser{Login: "John", Name: "Doe", PublicRepos: 7}
person.Name = "Bob"         // struct field


```

- Binary Operator: 

```go
* / % << >> & &^
```



### Statements
Statements are instructions that tell the Golang compiler what to do. Golang has statements for the following tasks:
- Declaration statements: Declaration statements are used to declare variables and constants.
```go
var int age  // variable declaration
const pi = 3.14 // constant declaration
type GithubUser struct { // type declaration 
	Login       string `json:"login"`
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

```
- Expression statements: Expression statements are used to evaluate expressions and return a value.
```go
a := 5
b := a * 2      // Expression statement (assignment)
fmt.Println(b)   // Expression statement (function call)

```
- Control flow statements: Control flow statements are used to control the order in which statements are executed.
```go
// Contidional statement
if x > 0 {
    // Code to execute if x is greater than 0
} else if x < 0 {
    // Code to execute if x is less than 0
} else {
    // Code to execute if x is equal to 0
}
```
Switch 
```go
switch day {
case "Monday":
    // Code for Monday
case "Tuesday":
    // Code for Tuesday
default:
    // Code for other days
}
```
Loop:
```go
for i := 0; i < 5; i++ {
    // Code to execute in a loop
}

// Infinite loop
for {
    // Code to execute indefinitely
}

for a, _: = range s {
}
```

- Selection statements: Selection statements are used to select a statement to execute based on a condition.
```go
select {
case <-ch1:
    // ...
case x := <-ch2:
    // ...use x...
case ch3 <- y:
    // ...
default:
    // ...
}
```
- Iteration statements: Iteration statements are used to execute a statement repeatedly.
- Jump statements: Jump statements are used to skip to a different statement.
- Assignment
```go
a : = 5            // a == 5
var b = 15         // b == 15 
c, d := 4, 8       // c == 4 and d == 8 
```
###  Comments
Comments are used to provide information about your code. Comments are not executed by the Golang compiler.: As you will have seen right up to now we have been heavily making use of comments. Comments are usually of two forms
- Line comments start with the character sequence // and stop at the end of the line. 
- General comments start with the character sequence /* and stop with the first subsequent character sequence */.
```go   
/* 1. example of mulitple line comment
   2. comment continues here
   3. and then it ends
*/
```



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
```
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


## Arrays/ Slices
As introduced in the data type section Array, Slice are reference types. 
An array in Go is a data structure that consists of an ordered sequence of elements of the same type.

Arrays in go are fixed size, meaning that the number of elements in the array cannot be changed once it is declared.
Arrays elements are referenced using indexes, The First element of an array is usually referenced with an index starting at 0.
Arrays are declared using the `[]` syntax, followed by the type of the elements and the number of elements in the array. For example,
if we have a set of integers of the following
`1, 2, 3, 4, 5, 6, 7,8,9,10`
 the following code declares an array of 10 integers as above:
```go
var numbers = [10]int{1,2,3,4,5,6,7,8,9,10}
```

while working with arrays, you are expected to know the index before you can get the elements. say we want to get the value 5 from the array above, we will need to know the index that 5 is in the array; in this case, the value 5 is in index 4. Below is how to retrieve 5
```go
var five int  = numbers[4]
```

It is a common practice to most at time go through the array. Suppose we want to scan an array moving from the first item to the last. There are various ways in which we can do that.  One simpler way if to use the range function to loop through the array. You will see more about this in later sections. 

```go
// Print the indices and elements.
for index, value := range numbers {
    fmt.Printf("%d %d\n", index, value)
}
```

Remember that we just discussed about functions before here. One thing to not is that just like you saw under the functions sections an array by default when passed to a function will be passed by value( a copy of the array is sent to the function), however there is a way in which we can ensure that an array is not passed by value but by reference as a reference type.  In order to archieve passing an array to a function by reference we use what is called pointer(*) you will learn more about pointers in subsequent section. 
```go
// function that takes an array of 10 integers and squares each one of the numbers. notice that the original array refered to by ptr will be modified inplace
func squareNumbers(ptr *[10]int) {
	for i := range ptr {
		ptr[i] = ptr[i] * ptr[i]
	}
}

// the caller can simply do the below to pass the array by reference
squareNumbers(&numbers)

```

## Slice
A slice is a data structure similar to an array but with a variable size. Slices are declared using the `[]` syntax, followed by the type of the elements and the slice elements. For example, the following code declares a slice of integers:
```go
var ages []int
```

The ages slice is initially empty, but it can be used to store any number of integers. To add elements to a slice, you can use the append() function. For example, the following code adds the number 40 to the ages slice:

```go
ages = append(ages, 40)
```
Slices also have a length and capacity. The length of a slice is the number of elements that it currently contains. The capacity of a slice is the maximum number of elements that the slice can contain without having to reallocate memory. to get the length of a slice we use the built in `len` function. 


### Differences between Arrays and Slices

The main difference between arrays and slices is that arrays are fixed-size, while slices are variable-size. This means you cannot change the number of elements in an array after it is declared, but you can change the number of elements in a slice at any time.

Another difference between arrays and slices is that slices contain a pointer to the underlying array. This means that if you change the contents of a slice, the underlying array will also be changed.


## Control Flow/Conditionals
Control flow is the order in which statements in a program are executed. Golang has a variety of control flow statements that can be used to control the order of execution. we can ensure that any custom code that we write will only execute in pre-defined circumstances. Most of the time, we want our program to perform a different action based on some certain scenario, It would be great if we can programmatically solve for such a use case. That is, have Go make these decisions for us when it comes to performing different actions at the appropriate times. Conditionals come in handy when doing these. You can now go back to the Statements section of this document and re-read to understand.

The most common conditional control flow statements in Golang are:
- If statement: The if statement is used to execute a block of code if a condition is true.
- Else statement: The else statement is used to execute a block of code if a condition is false.
- For loop: The for loop is used to execute a block of code repeatedly.
- Switch statement: The switch statement is used to execute a block of code based on the value of an expression.
- Selection statements: Selection statements are used to select a statement to execute based on a condition


### Loop

​​for INITIALIZATION; CONDITION; INCREMENT { // Code to be repeated 
}
1. The initialisation part sets the first(initial) value for the counting this is done before the first iteration.
2. The condition part is defined to check when the iteration should stop. This is evaluated during every round of the iteration, and if it's true, the loop body executes 
3. Lastly, the increment is defined to specify how the counting should change
Here is an example of a for loop:
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

This code will print the numbers from 0 to 9.

1. The initialisation is the part  i := 0
2. condition i < 10
3. increment i++


***Let's break it down step by step:***

- The loop begins with i pointing to a value of 0
- i < 10 is then evaluated. Since 0 is less than 10, the code block code is executed. As such, 0 is printed to the console.
- i  then increments by 1 and is therefore now 1.
- i < 10 is re-evaluated. Since 1 is still under 10, the code block is executed again, and now 1 is printed to the console.
- i then increment by 1, and is now 2.
- (etc.)
An alternative way for writing the above snippet is 

```go
package main
import "fmt"

func main() {
    i := 0
    for i < 10 {
        fmt.Print(i, " ")
        i++
    }
}
```


### Ranges
We had earlier mentioned range when dealing with arrays. A range in Go is a special syntax used to iterate over the elements of a collection. Collections in Go can include arrays, slices, maps, and channels.
The range syntax is as follows:
```go
for key, value := range collection {
    // Do something with key and value
}
```
The key and value variables are the index and value of the current element in the collection. For example, if the collection is an array, the critical variable will be the index of the element in the array, and the value variable will be the element's value.

The range syntax can also be used to iterate over the keys of a map. In this case, the key variable will be the key of the current element in the map, and the value variable will be the value of the element.

```go
package main 
import "fmt"
func main() {

    var ages = [4]int{12,24,17,25 }
    // loop over ages array using range
	for index, age := range ages {
		fmt.Println(index, age)
	}
    // create a map of strings keys and values of integers more on maps later
	var persons = map[string] int {"John": 12, "Mary": 24, "Alice": 17, "Bob": 25}

    // loop through persons map and display the names and ages.
	for name, age := range persons {
		fmt.Println(name, age)
	}
}
```


### Switch


Switch
Here is an example of a switch statement:
Code snippet

```go
switch day {
case "Monday":
    fmt.Println("Monday")
case "Tuesday":
    fmt.Println("Tuesday")
case "Wednesday":
    fmt.Println("Wednesday")
case "Thursday":
    fmt.Println("Thursday")
case "Friday":
    fmt.Println("Friday")
case "Saturday":
    fmt.Println("Saturday")
case "Sunday":
    fmt.Println("Sunday")
}
```

This code will print the day's name of the week based on the value of the variable day.

Each of these control flow statements has its own advantages and disadvantages. The if statement is the most straightforward control flow statement and is easy to understand. The else statement can be used to handle cases where the condition in the if statement is false. The for loop is a powerful control flow statement that can execute a code block repeatedly. The switch statement is a versatile control flow statement that can be used to execute a block of code based on the value of an expression. Ranges are a powerful tool for iterating over collections in Go. They are easy to use and can iterate over arrays, slices, maps, and channels.


## Map

We have used maps in the range section but now we dive deeper into what they are. Maps are a built-in data structure in Go which allows you to store and retrieve key-value pairs. In order programming languages, They are known as associative arrays(PHP), dictionaries(Python) or hashmaps(Java). Maps provide a way to store unordered data collection, each element being identified by a unique key. maps are a robust data structure that can store and retrieve data quickly and efficiently. The map's keys must be unique, while the value can be of any type.

### Creating a map
To create a map, we use the make function or the map literal syntax to create it.
1. Using the make function. 
`m:= make(map[keyType]valueType)`
2. Using the literal syntax.
`m:= map[keyType]valueType{}`

In the above syntax definition, the keyType is the type of keys in the map and will always have to be unique, while the valueType is the type of the value and can be of any supported datatype in go.  Eg if we want to create a map with keys to be strings and values to be integers, we can use the syntax below
```go
// using make function to create a map
var list = make(map[string]int)
list["John"] = 12
list["Mary"] = 24
list["Alice"] = 17
list["Bob"] = 25

for name, age := range list {
    fmt.Println(name, age)
}

// using the map literal syntax for creating map
var persons = map[string] int {"John": 12, "Mary": 24, "Alice": 17, "Bob": 25}

for name, age := range persons {
    fmt.Println(name, age)
}

```
*Note:* The key type Key must be comparable using ==, so that the map can test whether a given key is equal to one already within it.
Though floating-point numbers are comparable, it's a bad idea to compare floats for equality hence they are not good as keys. Value can however be of any type desired.

In the above, you will notice that we have not only declared how to create maps but also added how to add, iterate and access the map's values. To add values to a map, we assign a value to a specified key. 

```go
list["Carol"] = 30
persons["Carol"] = 30
```

Also, to access the values of a map.  we use the key inside of the square bracket as below
```go
fmt.Println(list["Carol"]) // Outputs: 30
fmt.Println(persons["Bob"]) // Outputs: 25
```
Note:

Accessing a key in a map will return the zero value for the value type if a key is unavailable. This essentially means that at any point in time if you try acessing any key on a map you will get a value. With this in mind, there is an optional second return value from accessing a map that can be used to decide whether the key exists. Map iteration is done using range, as we saw when discussing the range keyword.

```go
value, exists := persons["Carol"]
if exists {
    fmt.Println("Carol's score:", value)
} else {
    fmt.Println("Carol's score doesn't exist")
}
```
Deleting from a map

When deleting from a map, we use the delete function  eg.
```go
delete(persons, "Carol")
```
In case we call delete on a key that is not present in the map. nothing happens generally termed no-op hence it's safe to call delete even on keys that don't exist. question is why will you want to do that?

#### Exercise

1. Create a map of countries to their capitals.
2. Create a map of colors to their hexadecimal values.
3. Create a map of words to their definitions.
4.  Write a function in Go that takes a map of student names and their scores as input and returns the average score.  Refer to functions section above for help.

Check  Miscellanous for more


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

## Interfaces

Interfaces are a fundamental concept of Go. They allow you to define a set of method signatures without having to specify the implementation details, In essence, for us to archive polymorphism in Go, we make use of interfaces. Interfaces enable us to write modular and tightly coupled code to a specific implementation.

### Defining and interface.
To define an interface, we use the type keyword followed by the interface name and then the set of methods signatures. 
Below is the signature of how an interface looks like
```go
type <name> interface {
    method1 return_type
    method2 return_type
    method3 return_type
}
// Example 
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

In the example above, Shape is the name of the interface and it specifies two methods. Area() and Perimeter(). Any type implementing Area and Perimeter methods will implicitly satisfy the Shape interface.
To implement an interface, A type must provide the implementation for all the methods specified in the interface. Note that there is no explicit declaration or inheritance when implementing interfaces. If a type defines all the methods of an interface, it automatically satisfies that interface. 
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
type Square struct {
    side float64
}
func (s Square) Area() float64 {
    return s.side * s.side
}
func (s Square) Perimeter() float64 {
    return 4 * s.side
}
```


In the example above, the Square is a struct that implements a Shape interface by providing implementation of both the `Area()` and the `Perimeter()` method
To use the interface we created above, once the type satisfies the interface, values of that type can be assigned to the variable of the interface type. This allows abstraction and writing more generic code. 
```go
package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}
func (s Square) Perimeter() float64 {
	return 4 * s.side
}
func main() {
	s := Square{side: 5}
	var perimeter, area float64 = s.Area(), s.Perimeter()
	fmt.Print("Area: ", area, " Perimeter: ", perimeter)

}

// Outputs: Area: 20 Perimeter: 25
```

# Exercise
As a continuation of the above, Define a method that compares the area of a square to that of a Rectangle. Hints you can let the Rectangle also implement the interface for Shape such that your compare can generically make use of the Shape type.

2. Create an interface for a type that can be sorted.
3. Create a struct that implements the interface you created in the previous exercise.
4. Write a function that sorts a slice of values that implement the interface you created.


Goroutines
Goroutines are lightweight concurrency constructs in Go that allow you to perform concurrent operations without creating and managing threads explicitly. A goroutine is a function or method that can be executed concurrently with other goroutines, achieving concurrency straightforwardly and efficiently. Just like threads in other programming languages and operating systems threads, goroutines are similar to threads. When a Go programme starts. the only goroutine is the one that calls the primary function hence it's the main goroutine. From here on,  new go routines are created with the go keyword. 
Syntactically a go statement is a normal function or method call prefixed with the `go` keyword. Go statement causes a function to be called in a newly created goroutine different from the current executing main goroutine. 

***Note that:*** the go statement completes immediately.

To create a goroutine in Go, you simply prefix a function or method call with the go keyword:
`go myFunction()`.
Below is an example which uses a goroutine to print numbers from 0 to 15  with a delay of 500 milliseconds between each print. and in the primary function, we create a goroutine, allowing the program to sleep for 5 seconds to execute concurrently with the main program.

```go
package main

import (
	"fmt"
	"time"
)
func printNumbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(i)
	}
}

func main() {
	go printNumbers()
	time.Sleep(5 * time.Second)

}
```

Another good example for your understanding is a fibonacci programe uses an inefficient algorithm to calculate the Fibonacci of 45. Knowing that this algorithm will take quite some time to produce the result, we want to display spinner progress while the program calculates the fib of 45.  for general formula for fibonacci of n is `fib(n) = fib(n-1) + fib(n-2)` hence if we are to calculate fib(45) it will first try to calculate fib(44) + fib(43) notice that the process is very recursive and we keep repeating the process all over even though we had calculated the fib(43) when calculating fib(45), we will still redo it all over when calculating fib(44) This is why we say this solution is inefficient. We however are not interested in optimizing this here but to see how goroutines can help use display a feedback while the calculation happens inthe background.  

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go progressIndicator(500 * time.Millisecond)
	const n = 45
	fibn := fib(n)
	fmt.Printf("Fibunacci of %d is: =  %d", n, fibn)

}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func progressIndicator(delay time.Duration) {

	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}

}

```
The spinner runs while the computation is being carried out which takes time. changing n to 50 means it takes longer to calculate the fibonacci hence the spinner spins for a longer time

### Channels
In Go, communication between goroutines is archived with the use of channels(ie sending data from and to another goroutine). Channels are typed conduits for sending and receiving values. Channels allow for safe communication and synchronisation between goroutines. In order wards. "If goroutines are the activities of a concurrent go programme, then channels are the connections between them" You can also see it as follows  "if goroutines are the nodes of a graph, then channels are the vertices of the graph"


### Creating a channel.
To create a channel, we use the built-in make function.  eg     `ch := make(chan int)`
Note the following about channels
- A channel refers to the data structure created by the make function. 
- The zero value for a channel is nil. 
- Channels are passed by reference when they are being passed to functions.
- Channels can be compared with ==, resulting in a boolean true or false. True in the case that both are referencing the same data structure. 
- A channel has two primary functions. send and receive

Channel sends transfers a value from one goroutine to another over the channel to another goroutine which is listening to receive from the same channel. The `<-` operator is used for both sending and receiving. Depending on which side the channel is on. 

1. For a send operation, the channel is on the left `ch <-`; 
2. For a receive operation, the channel is on the operator's right `<- ch`.

```go
package main
import "fmt"
func main() {
    // Create a channel
    ch := make(chan int)
    // Start a goroutine
    go func() {
        // Send a value to the channel
        ch <- 5
    }()
    // Read a value from the channel
    fmt.Println(<-ch)
}

```

There is another last operation which can be done on a channel. This is a close operation. 
close(ch) Calling close on a channel indicates no more values will ever be sent on this channel. 


```go
package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)
	fibs := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// sending the value of i to the numbers channel
			fmt.Printf(" Sending %d to numbers channel\n", i)
			numbers <- i
		}
		close(numbers)
	}()

	// start a go routine for receiving numbers from the numbers channel

	go func() {
		for x := range numbers {
			// calculate the fibonacci of the number received in the channel and resend it to another channel
			fibs <- fib(x)
		}
		close(fibs)
	}()
	// read and print all the values received from the fibs channel
	for val := range fibs {
		fmt.Println("received from channel fib: ", val)
	}
}
```
The fib function(omitted in the example above) is the same inefficient one we implemented earlier you can challenge yourself and try optimizing the function.  Also we have made use of anynymouse functions for our go routines. One way to play around with the code snippet above is to refactor it to functions that take input and output channels, making the program more readable and modular.  There's a lot more to channels than what we have discussed; however, we end here for now and we will come back to the depth of it later.

#### Exercise

1. Create a function that takes in a sender channel ch and sends integers from 1 to 15. 
2. Create another receiver function that will receive all the values sent on the channel ch declared earlier and prints them. 
3. Combine and run program above by invoking the function with the go routine. 
4. Write a program that calculates the sum of a large slice of integers concurrently using goroutines. Split the slice into smaller parts and assign each part to a goroutine. Communicate the results using channels and calculate the final sum.


# Errors  and Error Handling

It's worth noting that even the most well-written program will and can encounter errors and as such success is not usually a surety and guaranteed outcome. For instance, if your program performs an input or output it will have to assume the possibility of errors for writing or reading from the input or output sources. Hence in order to know why these error or failures happens need to write our programs with assumptions and to take control of errors when they occur. Even for the most reliable operation, we should be sure to detect and know why it fails if it does happen. This is where error handling comes in handy. We therefore infer that error handling is a very crucial aspect of writing reliable, efficient and maintainable go code.
Whenever we create functions for which we have as a possible outcome to be an error, it is paramount to ensure that the function returns two results one which is the actual result and another which represents the error. Simply put the error should always be the last value in the list of returns from any function that potentially throws an error. When dealing with functions that return error, it is great to understand that there are two possible values for the value of type error. It can be nil or non-nil with nil representing that no error occurred and any other value apart from nil means there was an error.
Handling errors in go is by use of control flow. Also, when errors occur there are usually several ways in which the error can be dealt with. This also usually depends on the operation that was being performed and or the context in which we are operating.

1. ***Propagate the error back to the caller:*** This is usually when the error happens within a function that was called by a parent function the caller function will have to forward back the error to the caller and then terminate.
```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


const githubAPIURL = "https://api.github.com/users/"

type GithubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}
func mains() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	username := "fanyui"
	user, err := FetchGithubUser(ctx, username)
    // checking of error which can be returned from the FetchGithubUser function
	if err != nil {
		fmt.Println("Errror ", err)
		return
	}
	fmt.Printf("Github user infor: \n Username: %s\n Name: %s\n Repositoris: %d\n", user.Login, user.Name, user.PublicRepos)
}

// function that returns a user or an error it it happens
func FetchGithubUser(ctx context.Context, username string) (*GithubUser, error) {
	url := githubAPIURL + username
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	var user GithubUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
```

2. ***Retry attempt***: The second way of dealing with errors will be to retry the operation that was being performed with certain criteria and limits in most cases you want to have a sentinel which limits how many times you retry and if it continues to fail you want to stop doing the retries.
```go
func FetchGitHubUserWithExponentialBackoff(username string) (*GithubUser, error) {
	timeout := 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		resp, err := http.Get(githubAPIURL + username)
		if err == nil {
			var user GithubUser
			json.NewDecoder(resp.Body).Decode(&user)
			return &user, nil
		}
		fmt.Printf("Server is not responding. retrying\n")
		time.Sleep(time.Second << uint(tries))
	}

	return nil, fmt.Errorf("Timeout failed to fetch %s in %s time", githubAPIURL+username, timeout)
}
```


3. ***Logging the error***. In some cases, it's sufficient to just log a message of the error and continue execution or terminate. These are when the error is not of paramount importance then we log and continue however if the context here is main and in the case of terminating we log and then end or exit the program in its entirety. These are typical when the control flow is checked at the level of main.


***Note.*** Immediately after checking for errors, you should always deal with the error before the success. If the failure of the app causes the function to terminate then the success code is not needed to run as such this brings up another point. You should not have to use else for error handling but rather continue the code at the same root level as the control flow which initiated the error check.



## Zero value
Note: You have been hearing us made mention of zero value through out the document particularlyw when we talked about types and data types.  Now we explain what a zero value is about. 
The zero value

When storage is allocated for a variable, either through a declaration or a call of new, or when a new value is created, either through a composite literal or a call of make, and no explicit initialization is provided, the variable or value is given a default value. Each element of such a variable or value is set to the zero value for its type: false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.

These two simple declarations are equivalent:
```go
var i int
var i int = 0
```

```go
type T struct { i int; f float64; next *T }
t := new(T)

```
the following holds:
```go
t.i == 0
t.f == 0.0
t.next == nil
```
The same would also be true after

```go
var t T
```

[Reference Zero value](https://go.dev/ref/spec#The_zero_value)



## Tips
### Defer keyword
Did you know as your function grow and become more complex, acquire more resources, it becomes difficult to handle and cleaning up of these becomes event more difficult and hard to maintain? Well, just like many other programming languages, go has a solution for you which is the defer keyword. Syntactically, this is just a regular function or method prefixed with the defer keyword.
```go
func mains() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel() // Deferred cancel function
	username := "fanyui"
	user, err := FetchGithubUser(ctx, username)
	if err != nil {
		fmt.Println("Errror ", err)
		return
	}
	fmt.Printf("Github user infor: \n Username: %s\n Name: %s\n Repositoris: %d\n", user.Login, user.Name, user.PublicRepos)
}
```

The deferred statement is often used with antagonistic operations to ensure that resources are released no matter how complex the flow is. Some example of this antagonistic operations will be. open and close of files , connect and disconnect, luck and unlock etc We have made ample used of the defer function in our previous tips and implementations. eg
```go
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock() // deferred mutex unluck function
	return m[key]
}
```

As you will see in the example above, best practice suggest that you if a deferred statement is to be for releasing of a resource its best placed just after acquiring the resource itself. One thing to note about defer statement is that the function and argument expressions are evaluated when the statement is executed, but the actual call is deferred until the (caller) function that contains the defer statement has finished. The "finished" here means either normally or abnormally completed. By Normally, we mean the caller function executing a return statement or falling off the end. Abnormally: the caller function panicking.many function calls can be deferred and they are executed in the reverse of the order in which they were deferred.

### Leverage Context for Better Control and Cancellation

Context is a very powerful tool for controlling and managing the lifecycle of an application in go. this is most specifically useful when dealign with concurrent operations. Context helps you in setting deadlines, cancellation and setting request scoped values across api boundaries. Context comes in handy when ensuring that applications can gracefully handle deadlines cancellations and propagation of values through out the flow.
```go
func mains() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	username := "fanyui"
	user, err := FetchGithubUser(ctx, username)
	if err != nil {
		fmt.Println("Errror ", err)
		return
	}
	fmt.Printf("Github user infor: \n Username: %s\n Name: %s\n Repositoris: %d\n", user.Login, user.Name, user.PublicRepos)
}
```

Here are few uses of the context can be used
1. Cancel Multiple Operations Simultaneously: Here suppose you have multiple concurrent goroutines running and one fails and because of this failure you want to cancel all the others, you can rely on context and efficiently pass this cancellation command and terminate all the other tasks.

2. API calls and timeouts. There are often times when you don't want to spend time in waiting for responses from external or third party api. context can be used to set timeouts which cancel when the api takes long to respond.This is helpful as it ensure application behaves according to and responsive when external services are slow.
3. Graceful shutdown: There are cases when the application is exiting and had spawn background tasks. One way of making sure these are coordinated and the app shuts down gracefully is to leverage context.

Hey there are a lot more to it than highlighted above. Notice we've not even discussed on how to apply context with value context.WithValue(), WithDeadline() and WithCancel().


### sync.WaitGroup

One of the strengths of go is in it concurrency. However with all these, you sill need to have flexibility in coordinating how these concurrent  operations are performed. sync.waitGroup is particularly an important synchronization primitive designed to wait for completion of task(goroutines to complete). It's also very crutial in situations where you want ot coordinate completion.
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go ProcessTask(i, &wg)
	}
	wg.Wait()
	fmt.Println(" Finished processing, Done waiting")
}
func ProcessTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// perform the actual and logic work here by
	fmt.Println("processing with id", id)
}
```

Here are some of the reasons why you will need sync.WaitGroup in go 
1. Coordinating Concurrent Execution: When multiple goroutines are ran, they execute concurrently and if the main program doesn't wait for these to complete, they may exit before completing the task they were performing. In this case, the sync.WaitGroup is useful in ensuring that the main program waits until all goroutines are completed.
2. Preventing Premature Program Exit. Just like explained above, the main program can exit prematurely and this can lead to unexpected result. here sync.WaitGroup ensures main doesn't prematurely terminates.
3. Preventing Premature Program Exit: For programs that share resources between goroutines, it's crucial to wait for them to complete before proceeding with further exection and or exiting. In this case, sync.WaitGroup helps in coordinating safety between shared resources.
4. Handling Asynchronous Operations: There are often cases where we are dealing with asynchronous operations and or callbacks. This means we are expected to wait for the asyncrhonous tasks to finish. sync.WaitGroup helps here by allowing us wait for the task completion.

Sumarily, sync.WaitGroup is a powerful tool for orchestrating concurrent operations, preventing premature program exits, and coordinating shared resources in a safe manner in Go.