package structs

import (
	"fmt"

	"whatIsThis.com/user"
)

func StructMain() {
	userFirstName := user.GetUserData("Please enter your first name: ")
	userLastName := user.GetUserData("Please enter your last name: ")
	userBirthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin, err := user.NewAdmin("test@example.com", "test123")

	if err != nil {
		fmt.Println(err)
		return
	}

	admin.OutputUserDetails()

	// interface usable
	appUser.OutputUserDetails()

	appUser.ClearUserName()

}
