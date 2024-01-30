---
layout: page
title: Statement
nav_order: 5
permalink: /syntax/statement/
parent: Syntax

---

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
