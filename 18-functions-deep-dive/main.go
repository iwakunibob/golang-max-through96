package main

import "fmt"

func main() {
	decades := []int{1, 10, 100, 1000}
	numbers, sum := sumTotal(1, 2, 3, 4, 5, 6)
	fmt.Printf("Sum of %v = %v\n", numbers, sum)
	// pass slice to variadic function
	_, totalDecades := sumTotal(decades...)
	fmt.Printf("Sum of %v = %v\n", decades, totalDecades)
}

// Can accept a list of values ...
func sumTotal(numbers ...int) ([]int, int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return numbers, sum
}
