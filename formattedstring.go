//Formatted strings

package main

import ("fmt")

func main(){
  var a int;
  fmt.Scan(&a);
  fmt.Print("value of a : ",a);
  fmt.Printf("value of %d",a);

  //sprintfname := "Deven"
  age := 25
  result := fmt.Sprintf("Name: %s, Age: %d", name, age)
  fmt.Println(result)
// Output: Name: Deven, Age: 25
  
  name := "Deven"
  age := 25
  result := fmt.Sprint("Name: ", name, ", Age: ", age)
  fmt.Println(result)
// Output: Name: Deven, Age: 25

  name := "Deven"
  age := 25
  result := fmt.Sprintln("Name:", name, "Age:", age)
  fmt.Print(result)
  // Output: Name: Deven Age: 25\n

  var name string
  var age int
  fmt.Print("Enter name and age: ")
  fmt.Scan(&name, &age)
  fmt.Println("Name:", name, "Age:", age)

  var name string
  var age int
  fmt.Print("Enter name and age: ")
  fmt.Scanf("%s %d", &name, &age)

  var name string
  fmt.Print("Enter your name: ")
  fmt.Scanln(&name)

  fmt.Print("Hello")
  fmt.Print("World")
  // Output: HelloWorld

  fmt.Println("Hello", "World")
  // Output: Hello World\n

  name := "Deven"
  age := 25
  fmt.Printf("Name: %s, Age: %d\n", name, age)
  // Output: Name: Deven, Age: 25







}
