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
