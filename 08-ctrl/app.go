package main

import "fmt"

func main() {
	var choice, balance, amount = 0, 0.0, 0.0
	fmt.Print("*** WELCOME to hOurBank ATM ***\n")
	for {
		fmt.Print(
			"Select one of the functions by number:\n",
			"[1] = Check Balance\n",
			"[2] = Deposit Money\n",
			"[3] = Withdraw Money\n",
			"[4] = Exit\n",
			"Your choice:\n")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Printf("Your current balance is %.2f\n", balance)
		} else if choice == 2 {
			fmt.Print("Enter the amount of your deposit: $")
			fmt.Scan(&amount)
			if amount < 0 {
				fmt.Println("You cannot deposit a negative amount.")
				continue
			}
			balance += amount
			fmt.Printf("After the $%.2f deposit your new balance = $%.2f\n", amount, balance)
		} else if choice == 3 {
			fmt.Print("Enter the amount of your withdraw: $")
			fmt.Scan(&amount)
			if amount < 0 {
				fmt.Println("You cannot withdraw a negative amount.")
				continue
			} else if amount > balance {
				fmt.Println("You cannot withdraw more than the", balance, "account balance.")
				continue
			}
			balance -= amount
			fmt.Printf("After the $%.2f withdraw your new balance = $%.2f\n", amount, balance)
		} else if choice == 4 {
			break
		} else {
			print("^ Invalid Entry\n")
			continue
		}
		fmt.Println("Would you like to make another transaction? [1]=Yes [0]=No")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Print("\033c")
			continue
		} else if choice == 0 {
			break
		} else {
			print("^ Invalid Entry\n")
			continue
		}
	}
	fmt.Println("Thank you for banking with us. Good bye")
}
