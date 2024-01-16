package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64
	fmt.Print("\nEnter your revenue: $ ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your expenses: $ ")
	fmt.Scan(&expenses)
	fmt.Print("Enter your taxRate: % ")
	fmt.Scan(&taxRate)
	earnings := revenue - expenses
	tax := earnings * taxRate / 100
	profit := earnings - tax
	fmt.Println("Earnings before tax: $", earnings)
	fmt.Println("Tax: $", tax)
	fmt.Println("Profit after tax $", profit)
	fmt.Print("Earnings/Profit ratio:", earnings/profit, "\n")
}
