package main

// we're learning struct

type messageToSendTest struct {
	phoneNumber int
	message     string
}

// Structs can be nested to represent more complex entities:

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	// ?
	if mToSend.sender.name != "" && mToSend.recipient.name != "" && mToSend.sender.number != 0 && mToSend.recipient.number != 0 {
		return true
	}
	return false
}

// anonymous struct

type car struct {
  brand string
  model string
  doors int
  mileage int
  // wheel is a field containing an anonymous struct
  wheel struct {
    radius int
    material string
  }
}

var myCar = car{
  brand:   "Rezvani",
  model:   "Vengeance",
  doors:   4,
  mileage: 35000,
  wheel: struct{radius int; material string}{
    radius:   35,
    material: "alloy",
  },
}


// embedded structs

type sender struct {
	rateLimit int
}

type userStruct struct {
	name   string
	number int
	sender
}
