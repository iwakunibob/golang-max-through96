package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64
	revenue = getUserNum("\nEnter your revenue: $ ")
	expenses = getUserNum("Enter your expenses: $ ")
	taxRate = getUserNum("Enter your taxRate: % ")
	earnings, tax, profit := calcFinancials(revenue, expenses, taxRate)
	fmt.Printf("Earnings before tax: $%.2f\n", earnings)
	fmt.Printf("Tax: $%.2f\n", tax)
	fmt.Printf("Profit after tax $%.2f\n", profit)
	fmt.Printf("Earnings/Profit ratio:%.2f\n\n", earnings/profit)
}

func getUserNum(prompt string) float64 {
	fmt.Print(prompt)
	var val float64
	fmt.Scan(&val)
	return val
}

func calcFinancials(rev, expense, taxR float64) (earn, tax, prof float64) {
	earn = rev - expense
	tax = earn * taxR / 100
	prof = earn - tax
	return
}
