package user

import (
	"errors"
	"fmt"
)

type User struct {
	Firstname string
}

func (u *User) GetUserOP() {
	fmt.Println(u.Firstname)
}

func NewUser(firstname string) (*User, error) {

	if firstname == "" {
		return nil, errors.New("Enter the firstname")
	}

	return &User{
		firstname,
	}, nil
}
