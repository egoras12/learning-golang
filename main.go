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

func main () {
	answer := reformat("Bethrand", formatterFunc)

	fmt.Println(answer)
}
