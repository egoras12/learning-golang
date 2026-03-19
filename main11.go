package main

import (
	"fmt"
)

func updateNamePointer(x *string) {
	*x = "Khalifa"
}

func main11 () {
	fmt.Println("This is the 10th file explaining how pointers and refernce work")

	// Pointers just point to another memory block

	name := "Bethrand"

	namePointer := &name // Basically this "&" sign points to the location holding the value Bethrand.

	fmt.Printf("This is the pointer to the variable 'name': %v \n", namePointer)

	// We can also dereference a pointer to get the value stored at that memory location
	
	fmt.Printf("This is the value stored at the memory location pointed to by 'namePointer' %v \n", 	*namePointer) // The "*" sign dereferences the pointer to get the value stored at that memory location


	updateNamePointer(namePointer)

	fmt.Printf("You can see that after updating the value using the pointer, the new value is %v \n", name)

	fmt.Println("Testing again...")
}