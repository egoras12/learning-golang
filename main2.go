package main

import "fmt"

func main2() {

	// Arrays

	var namesInArray [4]string = [4]string{"Yoshi", "Mario", "Luigi", "Bee"}
	
	str := fmt.Sprintf("The names in the array are %v, and they are %v", len(namesInArray), namesInArray)

	fmt.Println(str)


	// Slices. Makes use of Array under the hood
	slicedArray := []float32{78, 90, 78.11}

	slicedArray[2] = 78 // Reassigning a new value to the index position of "2" in the array

	slicedArray = append(slicedArray, 89, 90, 1000)

	fmt.Println("The sliced array is", slicedArray)


	// Ranges, get a range of element and store them in a new slice
	// [1:4], gets 1 to 4 but removes 4
	// [1:]. gets from 1 to the end  
	// [:3]. gets from the beginning till 3
	rangeOne := slicedArray[3:];
	rangeTwo := slicedArray[:3];

	stringAgain := fmt.Sprintf("This is the 1st range '%v', and this is the 2nd range '%v' ", rangeOne, rangeTwo)

	fmt.Println(stringAgain)
	
}