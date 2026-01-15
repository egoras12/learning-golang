package main

import "fmt"

func main() {
	var namesArray [4]string = [4]string{"Yoshi", "Luigi", "Mario", "Bee"}

	// fmt.Println("This is the length of the array is", len(namesArray), "and they are", namesArray);

	fmt.Printf("The length of the array is %v, and these are the people in the array %v \n", len(namesArray), namesArray)
}