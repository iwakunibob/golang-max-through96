package main

import "fmt"

type product struct {
	title       string
	price       float64
	description string
	quanity     int
}

func main() {
	hobbies := [3]string{"bicycling", "programming", "xc-skiing"}
	fmt.Println(hobbies) // 1.
	fmt.Println(hobbies[:1])
	fmt.Println(hobbies[1:3]) // 2.
	bestHobbies := hobbies[0:2]
	otherHobbies := hobbies[:2]
	fmt.Println(bestHobbies, "\n", otherHobbies)
	otherHobbies = hobbies[1:3]
	fmt.Println(otherHobbies)
	courseGoals := []string{"learn go", "make go state machine"}
	fmt.Println(courseGoals)
	courseGoals[1] = "GPIO control"
	courseGoals = append(courseGoals, "State Machine", "Thesis Stuff")
	fmt.Println(courseGoals)

	// Bonus
	parts := []product{
		{"wheel", 34.5, "rubber tire and rim", 5},
		{"lights", 10, "halogen lights", 3},
	}
	fmt.Println(parts)
	addpart := product{"motor", 123, "v4 engine", 3}
	parts = append(parts, addpart)
	fmt.Println(parts)
}
