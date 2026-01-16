package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main() {

	// greetings := "Greetings Bee"

	// string1 := strings.Contains(greetings, "ings") // returns true if the 2nd argument exists in the 1st argument

	// fmt.Println(string1)

	// greetings = strings.ReplaceAll(greetings, "Bee", "Bethrand")
	// fmt.Println(greetings)


	// string3 := strings.Split(greetings, " ")
	// fmt.Println(string3)


	ages := []float64{7,9,3,10,789,56,32,976, 67.1, 67.12}

	ages = append(ages, 10,11, 3)

	sort.Float64s(ages)

	fmt.Printf("This is the sorted array of %v \n", ages)

	fmt.Println(sort.SearchFloat64s(ages, 3))
}