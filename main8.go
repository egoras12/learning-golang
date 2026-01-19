package main

import ("fmt")

func main8 () {

	fmt.Println("This one is for importing functions from other files and utilizing them")

	// make sure when running main.go, include main7.go; asin "go run main.go main7.go"
	firstInitial, secondInitial := getInitialsImplementation("bethrand nnaemeka chisom");

	fmt.Println(firstInitial, secondInitial)
}