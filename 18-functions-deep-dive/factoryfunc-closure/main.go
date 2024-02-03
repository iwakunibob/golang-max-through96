package main

import (
	"fmt"
)

func main() {
	dice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(dice)
	double := createFactorXformer(2)
	triple := createFactorXformer(3)
	square := createPowerXformer(2)
	cube := createPowerXformer(3)
	factorial := createFactorialXformer()
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
func createFactorialXformer() func(int) int {
	return func(num int) int {
		result := 1
		for i := 1; i <= num; i++ {
			result = result * i
		}
		return result
	}
}
func createFactorXformer(factor int) func(int) int {
	// Closure says the num and factor are recreated in factory func
	return func(num int) int {
		return num * factor
	}
}
func createPowerXformer(exponent int) func(int) int {
	// Closure says the num and factor are recreated in factory func
	return func(num int) int {
		power := num
		for i := 1; i < exponent; i++ {
			power = power * num
		}
		return power
	}
}
