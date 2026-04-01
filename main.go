package main

import (
	"fmt"
)

func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":	
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"

	// all three of these cases will set creator to "A Steve"
	case "macOS":
		fallthrough
	case "Mac OS X":
		fallthrough
	case "mac":
		creator = "A Steve"

	default:
		creator = "Unknown"
	}
	return creator
}

func main()  {
	
	if pi := 3.14; pi > 44 {
		fmt.Println("The 1st value is greater trhan the value of pi")
	} else if pi >12 {
		fmt.Println("The 2nd value is still greater than the value of pi")
	} else {
		fmt.Println("This is just the last block being called")
	}

	// Case Statements

	fmt.Printf("This is the creator of the %s operating system: %s\n", "windows", getCreator("windows"))
}




