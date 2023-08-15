package main

import (
	"fmt"
	"time"
)
func printNumbers()  {
	for i := 0; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}
	
func main()  {
	go printNumbers()
	time.Sleep(5 * time.Second)
}