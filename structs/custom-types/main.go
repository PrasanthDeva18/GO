package main

import "fmt"

//customer types

//syntax -  type MyType OriginalType

type str string

func (text str) getOp() {
	fmt.Println((text))
}

func main() {
	var a str = "prasanth"

	a.getOp()
}
