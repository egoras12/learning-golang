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

func main() {

	if pi := 3.14; pi > 44 {
		fmt.Println("The 1st value is greater trhan the value of pi")
	} else if pi > 12 {
		fmt.Println("The 2nd value is still greater than the value of pi")
	} else {
		fmt.Println("This is just the last block being called")
	}

	fmt.Printf("This is the creator of the %s operating system: %s\n", "macOS", getCreator("macOS"))

	// Case Statements
	polyMeter := 10

	switch polyMeter {
	case 10:
		fmt.Printf("This is the correct value of polymeter %d\n", polyMeter)

	case 100:
		fmt.Printf("This is the correct value of polymeter %d\n", polyMeter)

	case 1000:
		fmt.Printf("This is the correct value of polymeter %d\n", polyMeter)
	}

	// the _ means we should ignore the 2nd return statement
	tierVal, _ := multipleReturns(10000)

	fmt.Println(tierVal)

	fmt.Println(earlyGuardClauseReturn(188))
}

func multipleReturns(tier int) (string, string) {

	switch tier {
	case 10:
		return "This is the lowest tier", "Tier 10"

	case 100:
		return "This is the 2nd lowest tier", "Tier 100"

	case 500:
		return "This is the 3rd tier and it's an okay tier", "Tier 500"

	case 1000:
		return "This is the last tier, and it's an elite tier", "Tier 1000"

	default:
		return "This is a default tier, and it's not a good tier", "It means that the tier you entered is not valid"
	}
}

// ealry guard clause return

func earlyGuardClauseReturn(num int) string {
	if num == 11 {
		return "fail"
	}
	if num == 22 {
		return "fail"
	}

	if num == 10 {
		return "success"
	}

	return "failed in a worser manner"
}
