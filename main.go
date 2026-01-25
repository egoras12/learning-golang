package main

import (
	"fmt"
)

func updatename(x string) {
	x = "Emeka"
}

func main() {

	name := "Bethrnad"

	updatename(name)
	fmt.Println(name)
}

// Well still trying to understand this pass by value thing, but from what I get is that when we pass a variable to a function, it creates a copy of that variable and any changes made to that variable inside the function do not affect the original variable outside the function. So in this case, when we call updateName(name), it creates a copy of name and assigns it to x. When we change x to "Khalifa", it does not change the original name variable in main, which remains "Mia". Hence, when we print name in main, it still outputs "Mia".
