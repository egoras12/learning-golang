package main

import "fmt"

func main5() {
	fmt.Println("This is the next main.go file and it talks about conditionals and loops")

	age := 77.88

	if age != 77.8 {
		fmt.Println("This code block is supposed to run")
	} else {
		fmt.Println("This code block ran because the other was wrong")
	}

	namesInArray := []string{"Bethrand", "Bee", "John", "Doe", "Smith", "Jane", "Dane"}

	for index, value := range namesInArray {
		if index == 1 {
			fmt.Printf("This is postition %v and we'll continue the loop, instead of coming down becasue of the 'continue' keyword \n", index)
			continue
		}
		if index == 4 {
			fmt.Printf("This is index %v and the value here is %v, but we're breaking out here and we'll stop looping through because of the 'break' keyword \n", index, value)
			break
		}

		fmt.Printf("This is index %v and the value is %v \n", index, value)
	}
}
