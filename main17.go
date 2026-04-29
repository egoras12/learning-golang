package main

import "fmt"

func splitEmail(email string) (string, string) {
	// { 
	// remvoed this {} because it was not needed, and it was causing the variables to be scoped only within the block, which is not what we want.
		username, domain := "", ""
	// }
	
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

func main17 () {
	fmt.Println(splitEmail("bethrand2019@gmail.com"))
}