package main

import (
	"fmt"
)

func main() {
	dice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(dice)
	double := factorFactoryFunc(2)
	triple := factorFactoryFunc(3)
	square := powerFactoryFunc(2)
	cube := powerFactoryFunc(3)
	// factorial := factorialXformer()
	fmt.Println(transformNums(&dice, double))
	fmt.Println(transformNums(&dice, triple))
	fmt.Println(transformNums(&dice, square))
	fmt.Println(transformNums(&dice, cube))
	fmt.Println(transformNums(&dice, factorial))
}
func transformNums(nums *[]int, transform func(int) int) []int {
	dNums := []int{}
	for _, val := range *nums {
		dNums = append(dNums, transform(val))
	}
	return dNums
}
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
func factorFactoryFunc(factor int) func(int) int {
	// Closure says the num and factor are recreated in factory func
	return func(num int) int {
		return num * factor
	}
}
func powerFactoryFunc(exponent int) func(int) int {
	// Closure says the num and factor are recreated in factory func
	return func(num int) int {
		power := num
		for i := 1; i < exponent; i++ {
			power = power * num
		}
		return power
	}
}
