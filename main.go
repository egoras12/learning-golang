package main

import (
	"fmt"
)

func main () {

	fmt.Println("This is for learning about maps... I hear it's equivalent to objects in JS")


	mainMap := map[string]string{}

	mainMap = map[string]string{
		"name": "Bethrand Nnaemeka",
		"age": "26",
		"DOB": "21/10/2000",
	}

	fmt.Println(mainMap)

	mainMap["country"] = "Nigeria"
	
	fmt.Println(mainMap)

	mainMap2 := map[string]interface{}{
	"name": "Bethrand Nnaemeka",
	"age": 26,
	"DOB": "21/10/2000",
	"married": false,
	"height": 5.9,
	}

	fmt.Println(mainMap2["age"])

	// looping through the maps

	for key, value := range mainMap2 {
		fmt.Printf("The key is %v and value is %v \n", key, value)
	}

	// Declaring arrays
	// var mainArray []string = []string{"Tissue", "Toilet-paper"}
	// fmt.Println(mainArray)

	// Declaring maps [objects kind of]
	// var mainMap3 map[string]interface{} = map[string]interface{}{
	// 	"main-1": "Hello",
	// 	"main-2": "Hello-2",
	// 	"main-3": "Hello-3",
	// }
	// fmt.Println(mainMap3)
}