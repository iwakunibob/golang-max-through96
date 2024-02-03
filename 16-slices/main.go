package main

import (
	"fmt"
)

func main() {
	arrays()
	dynamicArrays()
}

func dynamicArrays() {
	itemPrices := []float64{1.5, 2.75}
	fmt.Println(itemPrices)
	morePrices := append(itemPrices, 3.25, 4.5, 5)
	fmt.Println(itemPrices)
	fmt.Println(morePrices)
	morePrices[1] = 7777.77
	fmt.Println(itemPrices)
	morePrices = morePrices[1:] // delete 1st
	fmt.Println(morePrices)
	morePrices = morePrices[:3] // delete last
	fmt.Println(morePrices)
}
