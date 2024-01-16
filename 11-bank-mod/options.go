package main

import "fmt"

func describeOptions() {
	fmt.Print(
		"Select one of the functions by number:\n",
		"[1] = Check Balance\n",
		"[2] = Deposit Money\n",
		"[3] = Withdraw Money\n",
		"[4] = Exit\n",
		"Your choice:\n")
}
