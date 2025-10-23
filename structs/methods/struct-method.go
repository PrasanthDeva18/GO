package main
import "fmt"

type Usertype struct {
	firstname string
}

func (u Usertype) getOutput() {
	fmt.Println(u.firstname)
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

}


// func getUserOp(user Usertype){

// }

func getUserInput(prompttext string) string {
	fmt.Println(prompttext);
	var a string;
	fmt.Scan(&a);
	
	return a
}