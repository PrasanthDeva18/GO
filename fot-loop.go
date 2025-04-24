package main

import "fmt"

func main() {
  for i := 0; i< 2; i++ {
    fmt.Println(i)
    }

  // it will run until we break
  for {
    fmt.Println("Hello word");

    if(true){
    break;
    }
    
  }

   for {
    fmt.Println("Hello word");

    if(true){
    continue;
    }
    
  }
}
