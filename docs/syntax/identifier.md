---
layout: page
title: Identifier
nav_order: 3
permalink: /syntax/identifier/
parent: Syntax

---


### Identifiers
Identifiers are used to identify variables, functions, and other entities in Golang. Identifiers can be made up of letters, numbers, and underscores. Identifiers must start with a letter or underscore

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
