
package main

import (
	"fmt"
)

func main4() {
	
	fmt.Println("This is the next main go file that talks about loops")


	// This is saying while x is less than 5, keep looping
	x := -10
	for x < 5 {
		fmt.Println("The value of x is:", x)
		x++
	}

	// This is a for loop that has initialization, condition and post statement
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("The value of i is:", i)
	// }

	// This is a for range loop
	// for i := range 10 {
	// 	fmt.Println("The value of i is:", i)
	// }

	namesInArray := []string{"Bethrand", "Bee", "John", "Doe", "Smith", "Jane", "Dane"}
	
	for i := 0; i < len(namesInArray); i++ {

		fmt.Println("This is the current name", namesInArray[i])
	}

	for index, value := range namesInArray {
		// fmt.Printf("The index is %d and the value is %s \n", index, value)
		fmt.Println(value, "is at index", index)
	}

	//  you want to use the value but you don't want index
	for _, value := range namesInArray {
		fmt.Println("The value is going to be", value)
	}

}