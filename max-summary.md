# Udemy Go course
by Max Schwarzmueller  
[Github Repo](https://github.com/mschwarzmueller/go-complete-guide-resources)

``` go 
package main	// this is the main package
import "fmt"  	// Import the fmt package
func main() {	// program starts here in main function
	fmt.Print("Hello Bob!\n")
} 
```
```bash
go run app.go  # Compiles a program
go build  # Used to compile and build an executable file
```
<https://pkg.go.dev/std>

package main  // Main entry point of 
can have multiple go files in one package  
can only have one main function in a pacakage  
can can have multiple packages import into a module  
go mod init max/prg01  
Null Values in go are the default uninitialized values  
int=>0 float64=>0.0  string=>""  bool=>false 

## Variable Declaration
```go
var investment, divRate, infRate, taxRate, years float64 
var output string  // Declaration
var choice, balance, amount, yesOrNo = 0, 0.0, 0.0, ""
taxRate := 15
const PI = 3.1415  // Constant values can't be changed
fmt.Print("\nEnter your Investment Amount: $ ") // Prompt
fmt.Scan(&investment)  // Input
```
### Short form implied type with initialization
```go	
tax := dividends * taxRate / 100
futValue := investment * math.Pow(1+divRate/100, years)
```
## Formatted Output and Strings
```go
output += fmt.Sprintf("\nFor a Principle = $%.0f\n", investment)
fmt.Print(output)
```
## Functions
```go
	func calcFv(investment, years, divRate float64) float64 {
		return investment * math.Pow(1+divRate/100, years)
	}  // end of function definition

	calcFv(divRate, dividends, years) // function call
```
### Muliple return value function
```go
func calcFvRfvDiv(invest, years, divRate, infRate float64) (float64, float64, float64) {
  fv := calcFV(invest, years, divRate)
  rfv := calcRealFV(fv, years, infRate)
  div := fv - invest
  return fv, rfv, div
}

futValue, realFutValue, dividends := calcFvRfvDiv(investment, years, divRate, infRate)
```
#### Alternative Syntax
```go
func calcFvRfvDiv(invest, years, divRate, infRate float64) (fv, rfv, div float64) {
  fv = calcFV(invest, years, divRate)
  rfv = calcRealFV(fv, years, infRate)
  div = fv - invest
  // return fv, rfv, div  // Optionally can do this or simply return
  return
}
```
	
## Control Structures
```go
if choice == 1 {
  fmt.Printf("Your current balance is %.2f\n", balance)
} else if choice == 2 {
  fmt.Print("Enter the amount of your deposit: $")
  fmt.Scan(&amount)
  balance += amount
  fmt.Printf("After the $%.2f deposit your new balance = $%.2f\n", amount, balance)
} else {
  fmt.Println("Thank you for banking with us. Good bye")
}

// switch statements do not use break for a loop. must use return
switch yesOrNo {
case "y":
  fmt.Print("\033c")
  continue
case "n":
  fmt.Println("Thank you for banking with us. Good bye")
  return
default:
  print("^ Invalid Entry\n")
  continue
}

for someCondition {
  // do something ...
}
```
## File I/O and Errors

	fmt.Scan(&amount) // For single word or number input
	fmt.Scanln(&amount) // For single word or number input and new line cancels
	bufio.NewReader(os.Stdin)


```go
func writeBalanceToFile(balance float64) {
  balanceText := fmt.Sprint(balance)
  os.WriteFile(accountBalanceFile, \[]byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
  data, err := os.ReadFile(accountBalanceFile)
  if err != nil { // Error handler
    return 0, errors.New("Error = Failed to find BankBalance.txt file")
  }
  balanceText := string(data)
  balance, err := strconv.ParseFloat(balanceText, 64)
  if err != nil { // Error handler
    return 0, errors.New("Error = Failed to parse strored balance value")
  }
  return balance, nil
}
```
### Error handling using return or panic function
```go
var balance, err = getBalanceFromFile()
if err != nil {
  fmt.Println("ERROR: ", err, "\n---------")
}	

if err != nil {
  fmt.Println("ERROR: ", err, "\n---------")
  return  // Can be used to terminate program on error
}

  if err != nil {
  fmt.Println("ERROR: ", err, "\n---------")
  panic("Can't continue, sorry") // Can be used to terminate program on error
}
```
## Modules splitting out code

### Multiple go files can be run together sharing functions using 
go run app.go options.go
	
### Modules 
 1. Use package name on line one
 2. Put package file in a new subfolder of project with mod
 3. go mod init folder/subfolder
 4. import package 
 5. All shared func in module must start upper-case character to share 
 modulename.FunctionName
 6. 3rd party packages imported with 
 Go package discovery page https://pkg.go.dev/
	https://go.dev/doc/code#ImportingRemote
	go get url # downloads project and adds to your project
	import package as specified by url
	
## Pointers

Avoid unnecessary value copies
pass-by-value creates duplicates of values
pass-by-refererence accesses a value directly
func outputAge{&age) // No copy is created

Directly mutate values  
agePointer := &age // Address or age  
No return value is required to change outputAge
```go
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
		*totalAge = *totalAge - 18 // Pass by reference no return value
		fmt.Println(totalAge)
	}
```
## Structs
```go

type user struct {
	firstName  string
	lastName   string
	birthDate  string
	modifyDate time.Time
}

func main() {
	var appUser1, appUser2 user // declaration of structure
	getUserData(&appUser1)
  getUserData(&appUser2)
	outputUserData(&appUser1)
	outputUserData(&appUser2)
	// appUserBob = user{  // Most clear form
	// 	firstName:  userfirstName,
	// 	lastName:   userlastName,
	// 	birthDate:  userbirthdate,
	// 	createDate: time.Now(),
	// }
	// appUserTom := user{ // Alternative form := if all fields and order same
	// 	userfirstName,
	// 	userlastName,
	// 	userbirthdate,
	// 	time.Now(),
	// }
}

func getUserData(usr *user) {
	usr.firstName = getStringData("Please enter your first name: ")
	usr.lastName = getStringData("Please enter your last name: ")
	usr.birthDate = getStringData("Please enter your birthdate (MM/DD/YYYY): ")
	usr.modifyDate = time.Now()
}

func getStringData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func outputUserData(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
```

### Go has no inherentance unlike OOP classes

Custom types to encapsulate related fields and method functions

Struct embedding allows you to extend an existing struct types including fields and methods. Note Admin struct below 

Validation of constructor methods and getters setters

```go
package main

import (
	"fmt"
	"max/14structs/user"
)

func main() {
	appUser1, err := user.New("Robert", "Laurie", "08091960")
	if err != nil {
		fmt.Println(err)
		return
	}
	var appUser2 user.User // declaration of structure
	appUser1.DisplayUser()
	appUser2.EditUser()
	appUser2.DisplayUser()
	appUser1.ClearUserName()
	appUser1.DisplayUser()
	admin1 := user.NewAdmin("happy@guy.com", "1234p")
	admin1.DisplayUser()
}

```
--- module max/14structs
```go
package user

import (
	"errors"
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
}

type Admin struct {
	email    string
	password string
	// Adm      User
	User // Anonymus is better than above
}

func New(fName, lName, bDate string) (*User, error) {
	if fName == "" || lName == "" || bDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &User{
		firstName: fName,
		lastName:  lName,
		birthDate: bDate,
	}, nil
}
func NewAdmin(email, pwd string) Admin {
	return Admin{
		email:    email,
		password: pwd,
		User: User{
			firstName: "Jim",
			lastName:  "Monroe",
			birthDate: "1935",
		},
	}
}
func (u *User) EditUser() {
	u.firstName = getStringData("Please enter your first name: ")
	u.lastName = getStringData("Please enter your last name: ")
	u.birthDate = getStringData("Please enter your birthdate (MM/DD/YYYY): ")
}
func (u *User) DisplayUser() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
func (u *User) ClearUserName() { // Modifiers must use *User
	u.firstName = ""
	u.lastName = ""
}

func getStringData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value) // Scanln Ends the input
	return value
}
```
### Creating Custom Types and Adding Methods

custom type can be used for aliases of existing types to associate methods

```go
type str string  // Alias for string type to associate methods
func main() {
	var name str = "Bob"		// Declaration for custom type str
	name.log
}
func (text str) log() {
	fmt.Println(txt)
}
```
#### Practice: Getting User input

Note in the following example program that json file is saved using `meta-data`. 

```go
package main

import (
	"bufio"
	"fmt"
	"max/note/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Save note failed")
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%s ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	// fmt.Scanln(&value) // Can't use because sentence not word
	return text
}
```
The Note Module max/note
```go
ppackage note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
func (n Note) Display() {
	fmt.Printf("\nTitle:\n%s\n\nContent:\n%s\n\n", n.Title, n.Content)
}
func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "-")
	fileName = strings.ReplaceAll(fileName, ",", "-")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
```
## Interfaces and Generic Code

Allows you to link func methods that can accept different types  
All types must have the same common methods

Can also embed other interfaces

```go
type saver interface { // An interface is a contract
	Save() error // All types using it must have a Save method
} // Note error/nil is returned

type outputable interface { // Embedded interface can contain
	saver     // other multiple interfaces
	Display() // other multiple methods
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data) // Embed interface func
}
func saveData(data saver) error { // Interface func
	err := data.Save()
	if err != nil {
		return err
	}
	return nil
}

```
### The Any value is allowed type interface{} "any" is an alias
```go 
func printAnything(value any) {
	fmt.Println(value)
}
```
### 104. Type switching based on a value type
```go
func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println(value)
	}
	default:  // for any other type or can omit to ignore processing
}
```
### 105. The extracting type operator can be used instead of switch
switch is more compact but this is more flexible
```go
func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}
}
```
### 107 Generic Types 
Can be used to write functions that work for multiple types
```go
func main() {
	result := add(1, 2)
	fmt.Println(result)
}
func add[T int | float64 | string](a, b T) T {
	return a + b
}
```
## 113. Arrays
Arrays can not be changed in Go  
Indexes cannot be negative  
Arrays are pass by reference even with = 
len() describes length of a slice  
cap() describes capacity of a slice to right
```go
func main() {
	// This is am array declaration not expandable
	var productNames [3]string
	productNames = [3]string{"Book", "Novel"}
	productNames[2] = "Carpet"
	var qty [3]int = [3]int{1, 3, 4}
	fmt.Println(productNames, qty)

	// This is a shorter combined form
	carMake := [3]string{"vw", "gm", "ford"}
	fmt.Println("cars = ", carMake)
	prices := [4]float64{3.5, 7.5, 12, 20.0}
	fmt.Println("prices = ", prices)

	// Arrays access using slices are references
	featuredPrices := prices[1:]
	fmt.Println("featured = ", featuredPrices)
	featuredPrices[0] = 299.99
	fmt.Println("featured = ", featuredPrices)
	fmt.Println("prices = ", prices)
	highlightedPrices := featuredPrices[:2]
	fmt.Println("highlighted = ", highlightedPrices)
	fmt.Println("prices = ", prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	highlightedPrices[0] = -5.4
	fmt.Println("prices = ", prices)
	highlightedPrices = highlightedPrices[0:]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
```
### 114. Slices are dynamic arrays
Slices can be declared without a dimension  
Can append with append(slice, 1, 2, ...)
Cand delete with slice reassignment
```go
func main() {
	itemPrices := []float64{1.5, 2.75}
	fmt.Println(itemPrices)
	morePrices := append(itemPrices, 3.25, 4.5, 5)
	fmt.Println(itemPrices)
	fmt.Println(morePrices)
	morePrices[1] = 7777.77
	fmt.Println(itemPrices)
	morePrices = morePrices[1:] // delete 1st
	fmt.Println(morePrices)
	morePrices = morePrices[:3] // delete last
	fmt.Println(morePrices)
}
```
#### Excercise 5 Arrays of Structs
```go
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
```
#### 116. Unpacking operator  
... unpacks list items individually where needed   
## 118. Maps  
Can use any type for keys  
Associates a key to data  
Similar to arrays with keys instead of indexes  
for loop using range is useful for arrays, slices index:value pairing,  
and maps  key:value pairing, 
```go
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
```
## 126. Functions Deep Dive
func is a first class value  
functions can be parameters for other functions in go  
```go
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
```
### 127. Returning functions values
Alias can be created for long string of types  
Functions can be returned from functions using id only
```go
package main

import "fmt"

// Can make custom type for long list of types
type transformFn func(int) int

func main() {
	decades := []int{1, 10, 100, 1000}
	fmt.Println(decades)
	// pass function as argument no ()
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
```
### 128. Anonymus Functions
Feature that defines a function at one place
Function is defined and called at one place in code
Function has no identifier so cannot be called from anywhere except where defined
```go
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
```
### 129. Closures and Factory functions
Closure uses variable value locked in to anonymous func
```go
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
func factorial(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result = result * i
	}
	return result
}
```
### Recursive function 
Call themselves recursively
must have a if-return to end recursion
Above factorial function done recursively
```go
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
```
### Variadic function 
Has a variable number of arguments that are turned into a slice
arguments are represented by ... meaning   
collect all remaining arguments as parameters and put in slice
Can also be used to convert slice to list of arguments

```go
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

```



