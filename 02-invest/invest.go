package main

import (
	"fmt"
	"math"
)

func main() {
	var investAmt float64 = 1000     // declaration and assignment with type
	expectedRate, years := 5.5, 10.0 // simplified declare and initialization
	futureValue := investAmt * math.Pow(1+expectedRate/100, years)
	fmt.Println(futureValue)
}
