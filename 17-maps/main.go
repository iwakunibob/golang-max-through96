package main

import "fmt"

// Alias types
type floatMap map[string]float64

// method func definition
func (m floatMap) output() {
	fmt.Println(m)
}
func main() {
	elements := map[string]string{
		"Au": "Gold",
		"Ag": "Silver",
		"H":  "Hydrogen",
		"He": "Helium",
	}
	fmt.Println(elements)
	elements["C"] = "Carbon"
	elements["O"] = "Oxygen"
	elements["Li"] = "Lithium"
	fmt.Println(elements)
	delete(elements, "He")
	fmt.Println(elements)
	makeArray()
	makeMap()
}
func makeArray() {
	// make keyword predefines array length 2 and capicity 5
	userNames := make([]string, 2, 5)
	userNames[0] = "Totoro"
	userNames = append(userNames, "Mei", "Sasuki")
	fmt.Println(userNames)
	for i, val := range userNames {
		fmt.Printf("Name[%d] = %s\n", i, val)
	}
}
func makeMap() {
	// make keyword defines map initial length for efficiency
	// courseRatings := make(map[string]float64, 4)
	courseRatings := make(floatMap, 5) // Using alias
	courseRatings["go"] = 4.7
	courseRatings["vue"] = 4.8
	fmt.Println(courseRatings)
	courseRatings.output() // Same with alias used for method
	for i, val := range courseRatings {
		fmt.Printf("Key: %s = Value: %f\n", i, val)
	}
}
