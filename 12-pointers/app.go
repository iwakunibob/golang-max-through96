package main

import "fmt"

func main() {
	age := 32
	// var agePointer *int
	agePointer := &age
	fmt.Println("Age = ", age, "Address of Age = ", &age)
	fmt.Println("Age = ", *agePointer, "Address of Age = ", agePointer)
	modifyAgeAdultYears(agePointer)
	fmt.Println("Number of years as and adult = ", age)
}

func modifyAgeAdultYears(totalAge *int) /*int */ {
	// return *totalAge - 18
	*totalAge = *totalAge - 18 // Pass by reference no return value required
	fmt.Println(totalAge)
}
