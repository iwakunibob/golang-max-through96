package main

import "fmt"

func main() {
	dice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(dice)
	// pass function as argument no ()
	doubled := multiplyNums(&dice, double)
	fmt.Println(doubled)
	tripled := multiplyNums(&dice, triple)
	fmt.Println(tripled)
}
func multiplyNums(num *[]int, multiple func(int) int) []int {
	mNums := []int{}
	for _, val := range *num {
		mNums = append(mNums, multiple(val))
	}
	return mNums
}
func double(num int) int {
	return num + num
}
func triple(num int) int {
	return num + num + num
}
