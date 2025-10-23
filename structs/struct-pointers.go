package main
import "fmt"

type Usertype struct {
	firstname string
}

func main(){

	ufirstname := getUserInput("enter the name");

	var user Usertype

	user = Usertype {
       ufirstname,
	}
	
	//pointers

	getUserOp(&user)
}


func getUserOp(user *Usertype){
    // (*u).firstname this method traditional way
    
	fmt.Println(user.firstname, (*u).firstname)
}

func getUserInput(prompttext string) string {
	fmt.Println(prompttext);
	var a string;
	fmt.Scan(&a);
	
	return a
}