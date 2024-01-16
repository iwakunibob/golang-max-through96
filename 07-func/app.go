package main

import (
	"fmt"
	"math"
)

func main() {
	var investment, divRate, infRate, taxRate, years float64
	var output string
	fmt.Print("\nEnter your Investment Amount: $ ")
	fmt.Scan(&investment)
	fmt.Print("Enter your timeframe in Years: ")
	fmt.Scan(&years)
	fmt.Print("Enter your Dividend Rate: % ")
	fmt.Scan(&divRate)
	fmt.Print("Enter your Inflation Rate: % ")
	fmt.Scan(&infRate)
	fmt.Print("Enter your Tax Rate: % ")
	fmt.Scan(&taxRate)
	// futValue := calcFV(investment, years, divRate)
	// realFutValue := calcRealFV(futValue, years, infRate)
	futValue, realFutValue, dividends := calcFvRfvDiv(investment, years, divRate, infRate)
	// dividends := calcDividend(futValue, investment)
	tax := dividends * taxRate / 100
	output = fmt.Sprintf("\nFor a Principle = $%.0f held for %v years,\n", investment, years)
	output += fmt.Sprintf(
		"with a Dividend Rate of %.1f%% the total Dividend = $%.2f, and the Future Value = $%.2f\n",
		divRate, dividends, futValue)
	output += fmt.Sprintf(
		"With an Inflation rate = %.1f%% the Real Future Value in todays money = $%.2f\n",
		infRate, realFutValue)
	output += fmt.Sprintf(
		"With a Tax rate = %.1f%% applied to total Dividends = $%.2f the Tax = $%.2f\n",
		taxRate, dividends, tax)
	output += fmt.Sprintf(
		"After %v years the initial Investment = $%.2f has an after Tax Real Future Value = $%.2f\n",
		years, investment, realFutValue-tax)
	output += fmt.Sprintf("In today's money. Spend it and be happy today :)\n\n")
	fmt.Print(output)
}

func calcFV(investment, years, divRate float64) float64 {
	return investment * math.Pow(1+divRate/100, years)
}

func calcRealFV(fv, years, infRate float64) float64 {
	return fv / math.Pow(1+(infRate)/100, years)
}

func calcDividend(futValue, investment float64) float64 {
	return (futValue - investment)
}

// func calcFvRfvDiv(invest, years, divRate, infRate float64) (float64, float64, float64) {
// 	fv := calcFV(invest, years, divRate)
// 	rfv := calcRealFV(fv, years, infRate)
// 	div := fv - invest
// 	return fv, rfv, div
// }

func calcFvRfvDiv(invest, years, divRate, infRate float64) (fv, rfv, div float64) {
	fv = calcFV(invest, years, divRate)
	rfv = calcRealFV(fv, years, infRate)
	div = fv - invest
	// return fv, rfv, div  // Optionally can do this or simply return
	return
}
