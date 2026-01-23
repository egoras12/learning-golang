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
