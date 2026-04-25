package main

import (
	"fmt"
)

// Teaches: anonymous functions can be created inline and passed as arguments.
func printReports(intro, body, outro string) {
	// Each inline function follows the same type: func(string) int.
	// It customizes cost logic for this specific report.
	printCostReport(func(message string) int {
		return 2 * len(message)
	}, intro)

	printCostReport(func(message string) int {
		return 3 * len(message)
	}, body)

	printCostReport(func(message string) int {
		return 4 * len(message)
	}, outro)
}

func main15() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

// Teaches higher-order functions: this function receives another function.
func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
