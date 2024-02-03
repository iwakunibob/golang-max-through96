package main

import "fmt"

// Can make custom type for long list of types
type transformFn func(int) int

func main() {
	dice := []int{1, 2, 3, 4, 5, 6}
	decades := []int{1, 10, 100, 1000}
	fmt.Println(dice)
	// pass function as argument no ()
	doubled := multipleOfNums(&dice, double)
	fmt.Println(doubled)
	tripled := multipleOfNums(&dice, triple)
	fmt.Println(tripled)
	transformFn1 := getTransformerFn(&decades)
	transformDecades := multipleOfNums(&decades, transformFn1)
	fmt.Println(transformDecades)
} //                                   func(int) int
func multipleOfNums(num *[]int, multiple transformFn) []int {
	mNums := []int{}
	for _, val := range *num {
		mNums = append(mNums, multiple(val))
	}
	return mNums
}
func getTransformerFn(sequence *[]int) transformFn {
	return double
}
func double(num int) int {
	return num + num
}
func triple(num int) int {
	return num + num + num
}
