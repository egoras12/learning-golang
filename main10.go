package main

import (
	"fmt"
)

func updateName(x string) {
	x = "Khalifa"
}

func updatePerson (p map[string]interface{}) {
	p["occupation"] = "Software Engineer"
}

func main10() {
	fmt.Println("This is the 10th file, explaining what Pass by Values are")

	// So, I learnt that a copy of the value of a variable is always made anytime a function is called.
	// The value passed into the functions argument is then stored in a new memory location.
	// For these data types, string, int, floats, bools, arrays, structs, when you pass them to functions, you are passing a copy of the value to the function.
	// Therefore, any changes made to the value inside the function does not affect the original value outside the function.
	// This is called Pass by Value.

	// But for these data types, slices, maps, channels, functions, when you pass them to functions, you are passing a copy of the reference to the value to the function.
	// Therefore, any changes made to the value inside the function affects the original value outside the function.
	// This is called Pass by Reference.

	name := "Emmanuel"
	fmt.Println(name)

	updateName(name)

	fmt.Printf("You can see that after logging the original value, then running the function before this line, we still get %v \n", name)

	// But let's try that for the second data types


	person := map[string]interface{}{
		"name": "Bethrand",
		"age": 22,
		"nationality": "Nigeria",
	}
	fmt.Println(person)

	updatePerson(person)

	fmt.Printf("But in this 2nd one, you can see that i logged the original 1st, ran the function afterwards and this data is being updated to this : %v \n", person)
}