package main

import ("fmt" 
"time")

type User struct {
	ufirstname string
	ulastname string
	ucreateAt time.Time
}

func main () {

	firstname := getInputUtils("Enter the first name");
	lastname := getInputUtils("Enter the last name");

	var appUser User

	// Struct instance, in that if you add struct key value pair in order key is no need

	// appUser = User{
	// firstname,
	//  lastname,
	//  time.Now(),
	// }

	// it will create a null values
	// appUser = User{}
	// if we omitted value set automatically it will set the null values
	appUser = User{
		ufirstname : firstname,
		ulastname : lastname,
		ucreateAt: time.Now(),
	}
	


	// fmt.Println(firstname, lastname)

	getUserData(appUser)
}

func getUserData(u User) {
	fmt.Println(u.ufirstname, u.ucreateAt)
}

func getInputUtils(prompttext string) string {
	fmt.Println(prompttext);

	var text string

	fmt.Scan(&text)

	return text;
}