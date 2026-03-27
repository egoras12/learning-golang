package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the next file on the line, and we're going to talk about struct in this file")
	fmt.Println("when running this one, input thi command ==> go run main12.go bill12.go. because 2 files are dependent on them")


	bill := bill12("Bethrand")
	bill13 := bill12("Beth") // this is a redeclaration of the bill variable, and it will cause an error. because we can't redeclare a variable in the same scope. to fix this, we can either change the name of the variable or we can use the short declaration operator (:=) to declare and initialize the variable in one step. like this: bill := bill12("Beth")

	fmt.Println(bill)
	fmt.Println(bill13)

	// Well, I learn about const today 

	const pi = 3.14 // this is a constant variable, and it cannot be changed. if we try to change the value of pi, it will cause an error. because constants are immutable, and they cannot be changed after they are declared.

	fmt.Println(pi)

	if pi > 4 {
		fmt.Println("Return null")
	} else {
		fmt.Println("Return pi")
	}
}
