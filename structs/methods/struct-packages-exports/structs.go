package main

import (
	"fmt"

	"example.com/struct-packages-exports/user"
)

func main() {

	fmt.Println("Learning Struct Constructor")

	var firstname string
	fmt.Scanln(&firstname)

	var appUserData *user.User

	appUserData, err := user.NewUser(firstname)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUserData.GetUserOP()

}
