package main

import (
	"fmt"
	"strings"
)

// First Implementation
func getInitialsImplementation(name string) (string, string) {

	name = strings.ToUpper(name)

	initials := strings.Split(name, " ")

	firstNameInitial := initials[0][:1]

	// Check length BEFORE accessing initials[1]
	if len(initials) > 1 {
		secondNameInitial := initials[1][:1]
		return firstNameInitial, secondNameInitial
	}

	return firstNameInitial, "_"
}

// Second Implementation
func getInitials(name string) (string, string) {
	// This function gets a name, turns it to uppercase, gets the initials and returns them
	name = strings.ToUpper(name)

	fullName := strings.Split(name, " ")

	initials := []string{}
	for _, value := range fullName {
		initials = append(initials, value[:1])
		// value[0] : this is byte
		// value[:1] : this is string
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {

	fmt.Println("This is the next main go file that talks about returning multiple values from functions")

	firstInitial, secondInitial := getInitialsImplementation("bethrand nnaemeka")
	fmt.Println("First and Second Initials:", firstInitial, secondInitial)

	thirdInitial, fourthInitial := getInitials("Luigi Raphael")
	fmt.Println("Third and Fourth Initials:", thirdInitial, fourthInitial)

	fifthInitial, sixthInitial := getInitialsImplementation("mario, luigi, bee")
	fmt.Println("Fifth and Sixth Initials:", fifthInitial, sixthInitial)

	seventhInitial, eighthInitial := getInitialsImplementation("mario")
	fmt.Println("Seventh and Eighth Initials:", seventhInitial, eighthInitial)
}
