package main

import (
	"fmt"

	users "example.com/embedding/user"
)

func main() {
	fmt.Println("Learning Embedding in Struct")
	// var appUser *users.User
	admin, err := users.NewUserAdmin("email@gmail.com", "1235")

	if err != nil {
		fmt.Println(err)
	}
	admin.GetUserOp()

}
