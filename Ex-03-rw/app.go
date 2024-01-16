package main

import (
	"errors"
	"fmt"
	"os"
)

const resultsFile = "ResultsFile.txt"

func main() {
	revenue, err1 := getUserNum("Revenue", "$")
	expenses, err2 := getUserNum("Expenses", "$")
	taxRate, err3 := getUserNum("Tax Rate", "%")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("ERROR:\n", err1, err2, err3)
		return
	}
	earnings, tax, profit := calcFinancials(revenue, expenses, taxRate)
	resultsFileOutput := fmt.Sprintf(
		"For Revenue=$ %.2f\nExpenses=$ %0.2f\nTax Rate=%0.1f%%\n---------\n",
		revenue, expenses, taxRate)
	resultsOutput := fmt.Sprintf("Earnings before tax: $%.2f\n", earnings)
	resultsOutput += fmt.Sprintf("Tax: $%.2f\n", tax)
	resultsOutput += fmt.Sprintf("Profit after tax $%.2f\n", profit)
	resultsOutput += fmt.Sprintf("Earnings/Profit ratio:%.3f", earnings/profit)
	fmt.Println(resultsOutput)
	resultsFileOutput += resultsOutput + "\n~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"
	os.WriteFile(resultsFile, []byte(resultsFileOutput), 0644)
}

func getUserNum(prompt, unit string) (float64, error) {
	var userInput float64
	promptMsg := fmt.Sprintf("Enter your %s: %s ", prompt, unit)

	fmt.Print(promptMsg)
	fmt.Scan(&userInput)
	if userInput < 0 {
		errorMsg := fmt.Sprintf("Value of %s must be a positive number\n", prompt)
		return 0, errors.New(errorMsg)
	}
	return userInput, nil
}

func calcFinancials(rev, expense, taxR float64) (earn, tax, prof float64) {
	earn = rev - expense
	tax = earn * taxR / 100
	prof = earn - tax
	return
}
