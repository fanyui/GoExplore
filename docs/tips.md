---
layout: page
title: Tips
permalink: /tips/
nav_order: 15

---
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