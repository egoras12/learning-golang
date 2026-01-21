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

	fmt.Println(mainMap2)
}