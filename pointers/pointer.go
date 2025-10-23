package main
import "fmt"

// we have worked on avoid duplicate of value in function

func main() {
	age := 20
	fmt.Println("age is", age)
		// agePointer := &age

	var agePointer *int;
	
	agePointer = &age
	
	fmt.Println(agePointer)

 // dereferencing

 fmt.Println(*agePointer)

	adultAge := getAdultYear(age)

	fmt.Println(adultAge)

	var val = 10;
	var a *int;
	a = &val;

	valOp := addition(a)

	valOp2 := addition(a)

	// valOp := addition(*val)

	fmt.Println(valOp)
	fmt.Println(valOp2)

}


func getAdultYear(age int) int {
	return age - 18;
}


func addition(a *int) int {
	// return a + 18

	*a = *a + 18
}


// using pointer for data mutation
func addition2(a *int)  {
	// return a + 18

	*a = *a + 18
}