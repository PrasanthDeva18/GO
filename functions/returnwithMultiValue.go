package main
import (
	"fmt"
)

func learn1()  {
   fmt.Println("hello world")
}

//named return value
func learn2 (b int, c string) (b1 int, c1 string) {
    b1 = b;
    c1 = c;
	return 
}

func learn3 (b int, c string) (int , string) {
	return b,c
}

func main() {
   learn1()
   a , b := learn2(5, "d")
   _ , d := learn2(5, "d")
   fmt.Println("hello",a,b );
   fmt.Println("hello",d );
}