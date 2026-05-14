package main

import "fmt"

func main25 () {
	user := newUser("Bethrand", "premium")
	message, sent := user.sendMessage("Hello, World!", 13)

	if sent {
		fmt.Printf("Message sent: %s\n", message)
	} else {
		fmt.Println("Message not sent: exceeds character limit.")
	}
}