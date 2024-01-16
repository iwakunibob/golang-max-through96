package main

import (
	"fmt"
	"math"
)

func main() {
	var investAmt, rateReturn, inflationRate, years float64
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investAmt)
	fmt.Print("Rate of Return %: ")
	fmt.Scan(&rateReturn)
	fmt.Print("Inflation Rate %: ")
	fmt.Scan(&inflationRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	futureValue := investAmt * math.Pow(1+rateReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
