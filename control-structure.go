pacakge main

import "fmt"

//banking application

func main(){
  let amountTotalBalance = 1000
  fmt.Println("Welcome to Deva bank");
  fmt.Println("Please select the option you want to proceed")
  fmt.Println("1. Check Balance");
  fmt.Println("2. Deposit");
  fmt.Println("3. Withdraw");
  fmt.Println("4. Reset Pin Number");
  
  var choice int;
  fmt.Println("Enter your Choice");
  fmt.Scan(&choice);

  if choice == 1 {
    fmt.Println("Available Balance", amountTotalBalance)
  } else if choice == 2 {
    var dA int;
    fmt.Println("Enter the amount to deposit);
    fmt.Scan(&dA);
    amountTotalBalance := amountTotalBalance +  dA
    fmt.Println("Availble Balance, amountTotalBalance);
  } else {
  fmt.Println("There is no choice");  
  }

}
