package users

import (
	"errors"
	"fmt"
)

type User struct {
	firstName string
	lastName  string
}

type Account struct {
	email    string
	password string
	User
}

func (U *Account) GetUserOp() {
	fmt.Println(U.firstName, U.email)
}

func NewUserAdmin(email, password string) (*Account, error) {
	if email == "" {
		return nil, errors.New("Please fill the fields")
	}

	return &Account{
		email:    email,
		password: password,
		User: User{
			firstName: "Prasanth",
			lastName:  "deva",
		},
	}, nil

}

func NewUser(firtname, lastname string) (*User, error) {
	if firtname == "" {
		return nil, errors.New("Please fill the fields")
	}

	return &User{
		firtname,
		lastname,
	}, nil

}
