package main

import (
	"11-bank-mod/fileops"
	"fmt"
)

const accountBalanceFile = "BankBalance.txt"

func main() {
	var choice, amount, yesOrNo = 0, 0.0, ""
	var balance, err = fileops.GetNumFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR: ", err, "\n---------")
		panic("Can't continue, sorry") // Can be used to terminate program on error
	}
	fmt.Print("*** WELCOME to hOurBank ATM ***\n")
	for {
		describeOptions()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("Your current balance is %.2f\n", balance)
		case 2:
			fmt.Print("Enter the amount of your deposit: $")
			fmt.Scan(&amount)
			if amount < 0 {
				fmt.Println("You cannot deposit a negative amount.")
				continue
			}
			balance += amount
			fmt.Printf("After the $%.2f deposit your new balance = $%.2f\n", amount, balance)
			fileops.WriteNumToFile(balance, accountBalanceFile)
		case 3:
			fmt.Print("Enter the amount of your withdraw: $")
			fmt.Scan(&amount)
			if amount < 0 {
				fmt.Println("You cannot withdraw a negative amount.")
				continue
			} else if amount > balance {
				fmt.Println("You cannot withdraw more than the $", balance, "account balance.")
				continue
			}
			balance -= amount
			fmt.Printf("After the $%.2f withdraw your new balance = $%.2f\n", amount, balance)
			fileops.WriteNumToFile(balance, accountBalanceFile)
		case 4:
			fmt.Println("Thank you for banking with us. Good bye")
			return
		default:
			print("^ Invalid Entry\n")
		}
		fmt.Println("Would you like to make another transaction? [y]=Yes [n]=No")
		fmt.Scan(&yesOrNo)
		switch yesOrNo {
		case "y":
			fmt.Print("\033c")
			continue
		case "n":
			fmt.Println("Thank you for banking with us. Good bye")
			return
		default:
			print("^ Invalid Entry\n")
			continue
		}
	}
}
