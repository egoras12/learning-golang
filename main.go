package main

import (
	"fmt"
)

func updateName(x string) {
	x = "Khalifa"
}

func main() {
	fmt.Println("This is the 10th file, explaining what Pass by Values are")

	name := "Mia"

	updateName(name)

	fmt.Println(name)
}

// Well still trying to understand this pass by value thing, but from what I get is that when we pass a variable to a function, it creates a copy of that variable and any changes made to that variable inside the function do not affect the original variable outside the function. So in this case, when we call updateName(name), it creates a copy of name and assigns it to x. When we change x to "Khalifa", it does not change the original name variable in main, which remains "Mia". Hence, when we print name in main, it still outputs "Mia".