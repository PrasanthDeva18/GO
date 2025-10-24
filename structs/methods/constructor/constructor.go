package main

import (
	"errors"
	"fmt"
)

type User struct {
	firstname string
}

func (u *User) getUserOP() {
	fmt.Println(u.firstname)
}

func newUser(firstname string) (*User, error) {

	if firstname == "" {
		return nil, errors.New("Enter the firstname")
	}

	return &User{
		firstname,
	}, nil
}

func main() {

	fmt.Println("Learning Struct Constructor")

	var firstname string
	fmt.Scanln(&firstname)

	var appUserData *User

	appUserData, err := newUser(firstname)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUserData.getUserOP()

}
