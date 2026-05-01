package main

import "fmt"

func adder() func(int) int {
	// WE're going to learn closures in this function. A closure is a function that captures the variables in its surrounding scope. In this case, we want to capture the variable "sum" so that we can keep track of the running total.

	var sum int
	sum = 0

	return func (x int) int {
		sum += x
		return sum
	}
}


func main19 () {

	add := adder()

	add(1)
	add(2)
	add(3)

	fmt.Println(add(4))
}