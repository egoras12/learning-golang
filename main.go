package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the next file on the line, and we're going to talk about struct in this file")
	fmt.Println("when running this one, input thi command ==> go run main12.go bill12.go. because 2 files are dependent on them")


	bill := bill12("Bethrand")

	fmt.Println(bill)
}