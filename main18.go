package main

import (
	"fmt"
)

func main18() {
	checkOfOrderReq, accountBalance := placeOrder("2", 8, 20)

	if checkOfOrderReq {
		fmt.Printf("The order req was accepted as %v and your available balance is %v \n", checkOfOrderReq, accountBalance)
	} else {
		fmt.Printf("The order req was accepted as %v and your available balance is %v. The reason could be because of insufficient balance or invalid productID \n", checkOfOrderReq, accountBalance)
	}
}

func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	// Rule 1: reject unknown products.
	// Here, price <= 0 is treated as an invalid product ID.

	priceOfProduct := priceList(productID)

	if priceOfProduct <= 0.00 {
		return false, accountBalance
	}

	// Rule 2: reject if requested quantity exceeds stock.

	availableStockForProduct := amountInStock(productID)

	if quantity > availableStockForProduct {
		return false, accountBalance
	}

	// Rule 3: reject if user cannot afford total order price.
	// Optimization idea (no code change): you already have priceOfProduct,
	// so you could compute total from it to avoid looking up product price again.
	priceOfOrder := calcPrice(productID, quantity)

	if priceOfOrder > accountBalance {
		return false, accountBalance
	}
	// Success path: return true and the updated balance.
	// Assignment detail: on failures above, you still return the original balance.

	newBalance := accountBalance - priceOfOrder
	return true, newBalance

}

// Don't touch below this line

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	} else if productID == "2" {
		return 2.25
	} else if productID == "3" {
		return 3.00
	} else if productID == "4" {
		return 1.00
	} else if productID == "5" {
		return 2.50
	} else if productID == "6" {
		return 8.99
	} else if productID == "7" {
		return 22.50
	} else if productID == "8" {
		return 50.00
	} else if productID == "9" {
		return 999.99
	} else {
		return 0.00
	}
}

func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	} else if productID == "2" {
		return 25
	} else if productID == "3" {
		return 4
	} else if productID == "4" {
		return 6
	} else if productID == "5" {
		return 50
	} else if productID == "6" {
		return 2
	} else if productID == "7" {
		return 0
	} else if productID == "8" {
		return 99
	} else if productID == "9" {
		return 1
	} else {
		return 0
	}
}
