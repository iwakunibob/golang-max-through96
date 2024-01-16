package main

import (
	"fmt"
	"math"
)

func main() {
	var investment, divRate, infRate, taxRate, years float64
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
	futValue := investment * math.Pow(1+divRate/100, years)
	realFutValue := futValue / math.Pow(1+(infRate)/100, years)
	dividends := futValue - investment
	tax := dividends * taxRate / 100
	fmt.Printf("\nFor a Principle = $%.0f held for %v years,\n", investment, years)
	fmt.Printf("with a Dividend Rate of %.1f%% the total Dividend = $%.2f, and the Future Value = $%.2f.\n", divRate, dividends, futValue)
	fmt.Printf("With an Inflation rate = %.1f%% the Real Future Value in todays money = $%.2f\n", infRate, realFutValue)
	fmt.Printf("With a Tax rate = %.1f%% applied to total Dividends = $%.2f the Tax = $%.2f\n", taxRate, dividends, tax)
	fmt.Printf("After %v years the initial Investment = $%.2f has an after Tax Real Future Value = $%.2f\n", years, investment, realFutValue-tax)
	fmt.Printf("In today's money. Spend it and be happy today :)\n\n")
}
