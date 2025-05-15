package main
import (
  "fmt"
  "panic"
  "errors"
)

func errorfunction()(error){
  return errors.New("this is from error function");
}

// panic keyword usage
func main(){
    var err = errorfunction()
    if err != nil {
        fmt.Println(err);
        // return 
        // panic("can't continue the code") 
      panic(err) 
    }
}
