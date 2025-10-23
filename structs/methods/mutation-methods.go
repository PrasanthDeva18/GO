package main
import "fmt"

type Usertype struct {
	firstname string
}

func (u *Usertype) getOutput() {
if u.firstname != ""  {
fmt.Println(u.firstname)
	} else {
		fmt.Println("first name is cleared")
	}
}

func (u *Usertype) clearValue() {
	// fmt.Println(u.firstname)
	u.firstname = ""
}

func main(){

	ufirstname := getUserInput("enter the name");

	var user Usertype

	user = Usertype {
       ufirstname,
	}
	
	//pointers

	// getUserOp(&user)
	user.getOutput()
	user.clearValue()
		user.getOutput()

}


// func getUserOp(user Usertype){

// }

func getUserInput(prompttext string) string {
	fmt.Println(prompttext);
	var a string;
	fmt.Scan(&a);
	
	return a
}