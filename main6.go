package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning, %v! \n", name)
}

func sayGoodBye(name string) {
	fmt.Printf("Good bye, %v! \n", name)
}

func sayMultipleGreetings(name []string, function1 func(string), function2 func(string)) {
	for _, value := range name {
		function1(value)

		function2(value)
	}

}

func areaOfCircle(radius float64) float64 {
	return math.Pi * (radius * radius)
}

func main6() {
	fmt.Println("This is the next main.go file, and this is for functions")

	sayMultipleGreetings([]string{"Luigi", "Mario", "Bee"}, sayGreeting, sayGoodBye)

	area := areaOfCircle(10.5)
	fmt.Println(area)

	// Print in 3 decimal places
	fmt.Printf("The area of the circle is %0.3f", area)
}
