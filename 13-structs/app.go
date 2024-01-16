package main

import (
	"fmt"
	"time"
)

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
