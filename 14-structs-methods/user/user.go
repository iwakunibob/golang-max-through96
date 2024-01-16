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
