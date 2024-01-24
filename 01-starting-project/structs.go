package main

import (
	"fmt"

	"example.com/structs/user"
)

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("bob@aol.com", "password")

	admin.OutputUserDetails()
	admin.User.ClearUsername()
	admin.User.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()

	var test customString = "this is a test"

	test.log()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
