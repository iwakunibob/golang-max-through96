package main

import "fmt"

func arrays() {
	// This is am array declaration not expandable
	var productNames [3]string
	productNames = [3]string{"Book", "Novel"}
	productNames[2] = "Carpet"
	var qty [3]int = [3]int{1, 3, 4}
	fmt.Println(productNames, qty)

	// This is a shorter combined form
	carMake := [3]string{"vw", "gm", "ford"}
	fmt.Println("cars = ", carMake)
	prices := [4]float64{3.5, 7.5, 12, 20.0}
	fmt.Println("prices = ", prices)

	// Arrays access using slices are references
	featuredPrices := prices[1:]
	fmt.Println("featured = ", featuredPrices)
	featuredPrices[0] = 299.99
	fmt.Println("featured = ", featuredPrices)
	fmt.Println("prices = ", prices)
	highlightedPrices := featuredPrices[:2]
	fmt.Println("highlighted = ", highlightedPrices)
	fmt.Println("prices = ", prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	highlightedPrices[0] = -5.4
	fmt.Println("prices = ", prices)
	highlightedPrices = highlightedPrices[0:]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
