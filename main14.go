package main

import (
	"fmt"
)

func formatterFunc(message string) string {
	return "TEXTIO: " + message + "..."
}

func reformat(message string, formatter func(string) string) string {
	return formatter(message); 
}

func main14 () {
	answer := reformat("Bethrand", formatterFunc)

	fmt.Println(answer)

	stri := fmt.Sprintln("Did not learn any thing today, and that's a bummer")
	fmt.Println(stri)
}
