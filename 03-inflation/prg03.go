package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5 // Can't be changed
	investAmt := 1000.0
	expectedRate := 5.5
	years := 10.0
	futureValue := investAmt * math.Pow(1+expectedRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
