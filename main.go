package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// Let's clean up Textio's authentication logic. We store our user's authentication data inside an authenticationInfo struct. We need a method that can take that data and return a basic authorization string.

func (auth authenticationInfo) login() string {
	return fmt.Sprintf("Authorization: Basic %v:%v", auth.username, auth.password)
}

func main() {

	person := authenticationInfo{
		"Bethrand",
		"1234567890",
	}

	fmt.Println(person.login())
}
