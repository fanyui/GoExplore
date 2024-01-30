---
layout: page
title: Programe structure
permalink: /program-structure/
nav_order: 3
---

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

This marks the end of our first program in Go. Excited?. There's more to come. 
