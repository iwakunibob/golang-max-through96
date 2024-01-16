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
