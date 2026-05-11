// package main                      // first implementation, before refactoring

// import "fmt"

// type User struct {
// 	Name string
// 	Membership
// }

// type Membership struct {
// 	Type string
// 	MessageCharLimit int
// }

// func newUser(name string, membershipType string) User {
// 	// ?

// 	if membershipType == "premium" {
// 		return User{
// 			Name: name,
// 			Membership: Membership{
// 				Type: "premium",
// 				MessageCharLimit: 1000,
// 			},
// 		}
// 	} else {
// 		return User{
// 			Name: name,
// 			Membership: Membership{
// 				Type: "regular",
// 				MessageCharLimit: 100,
// 			},
// 		}
// 	}
// }

// func main24 () {
// 	fmt.Println(newUser("Bethrand", "premium"))
// }

package main

import "fmt"

//

func (u User) sendMessage(message string, messageLength int) (string, bool) {

	if messageLength > u.MessageCharLimit {
		return "", false
	}

	return message, true
}

// don't touch below this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}

func main() {

	// first implementation, before refactoring

	// newUserMember := User{
	// 	"Bethrand",
	// 	Membership{
	// 		Type:             "premium",
	// 		MessageCharLimit: 1000,
	// 	},
	// }

	// fmt.Println(newUserMember.sendMessage("Hello, World!", len("Hello, World!")))

	user := newUser("Bethrand", "premium")

	msg, ok := user.sendMessage("Hello World", len("Hello World!"))
	fmt.Println(msg, ok)
}
