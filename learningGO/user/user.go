package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

//used this file with structs/ folder

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// user constructor
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First/Last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// admin constructor
func NewAdmin(email, password string) (*Admin, error) {
	if !govalidator.IsEmail(email) {
		return nil, errors.New("please write a valid email")
	} else if password == "" {
		return nil, errors.New("please write a password")
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}, nil
}

// to get User data
func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// prints all data
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// clear username
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
