package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate, tax float64
	fmt.Println("This program will calculate the Profit based on\n",
		"entered revenue, expenses, and tax rate.")
	fmt.Print("Enter Revenue $: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses $: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter Tax Rate %: ")
	fmt.Scan(&taxRate)
	earnings := revenue - expenses
	fmt.Println("\nEarnings before Tax = $", earnings)
	tax = earnings * taxRate / 100
	fmt.Println("Tax = $", tax)
	profit := earnings - tax
	fmt.Println("Profit after Tax = $", profit)
	fmt.Println("Earnings / Profit Ratio = ", earnings/profit)
}
