package main

import "fmt"

func main() {
	dice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(dice)
	doubleRoll := transformNumbers(&dice, func(num int) int {
		return num * 2
	})
	fmt.Println(doubleRoll)
	tripleRoll := transformNumbers(&dice, func(num int) int {
		return num * 3
	})
	fmt.Println(tripleRoll)
	squareDice := transformNumbers(&dice, func(num int) int {
		return num * num
	})
	fmt.Println(squareDice)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNums := []int{}
	for _, val := range *numbers {
		dNums = append(dNums, transform(val))
	}
	return dNums
}
