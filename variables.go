package main
import "fmt"

func main(){
  // types of variable
  // float64, string, int, bool
  // default assign of value float64 -> 0.0, int -> 0, string -> "", bool -> false
  
  var a = 10; //type inferred by go
  var b string = "Prasanth";
  c := 20; //type inferred by go
  var temp1, temp2, temp3 int = 10, 20, 30; 
  var type1, type2, type3 = 40, 50, 60; // type infered by go
  other1, other2, other3 := 10, 20, 30, 40;

  //variable declaration
  var types1 string;
  var types2 int;
  
  fmt.Println(b);
  fmt.Print(c);
}
