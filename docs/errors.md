---
layout: page
title: Errors  and Error Handling
permalink: /errors/
nav_order: 13

---

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

