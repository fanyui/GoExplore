---
layout: page
title: Control flow
permalink: /control-flow/
nav_order: 9

---



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
