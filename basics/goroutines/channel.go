package main

import "fmt"

func main() {
	// Create a channel
	numbers := make(chan int)
	fibs := make(chan int)
	// Start a goroutine
	go func() {
		for i := 0; i < 10; i++ {
			// Send a value to the channel
			fmt.Printf(("Sending %d to numbers channel\n"), i)
			numbers <- i
		}
		close(numbers)
	}()
	// Start a goroutine for receving values from the numbers channel, calculating fib and then sending them to the fibs channel
	go func() {
		for x := range numbers {
			fibs <- fib(x)
		}
		close(fibs)
	}()			
	// Read a value received from the fib channel
	for x := range fibs {
		fmt.Printf("Received %d from fibs channel\n", x)
	}
}
// calculate fibonacci number
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
