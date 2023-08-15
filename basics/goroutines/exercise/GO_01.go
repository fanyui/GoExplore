package main

import "fmt"

func main() {
	ch := make(chan int)
 go sender(ch)
 receiver(ch)

}
func sender(ch chan<- int) {
	// send from 1 to 15
	for i := 1; i <= 15; i++ {
		ch <- i
	}
	close(ch)
}
func receiver(ch <-chan int) {
	// receive 15 times
	for i:= range ch {
		fmt.Println(i)
	}
}