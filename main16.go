package main

import (
	"fmt"
)

func bootup() {
	// Learned: defer runs when this function returns, no matter which path returns.
	defer fmt.Println("TEXTIO BOOTUP DONE")
	// Learned: one defer avoids repeating the same print before every return.
	ok := connectToDB()
	if !ok {
		// Early return still triggers deferred print.
		return
	}
	ok = connectToPaymentProvider()
	if !ok {
		// Same here: deferred print runs on this return too.
		return
	}
	// Normal success path, then deferred print runs at function exit.
	fmt.Println("All systems ready!")
}

// Deferred code runs on both failure paths and the success path

var shouldConnectToDB = true

func connectToDB() bool {
	fmt.Println("Connecting to database...")
	if shouldConnectToDB {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

var shouldConnectToPaymentProvider = true

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

func testProject(dbSuccess, paymentSuccess bool) {
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("====================================")
}

func main16() {
	testProject(true, true)
	testProject(false, true)
	testProject(true, false)
	testProject(false, false)
}
