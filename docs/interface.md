---
layout: page
title: Interface
permalink: /interface/
nav_order: 12

---

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

