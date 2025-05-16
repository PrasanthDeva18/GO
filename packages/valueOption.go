package main

import "fmt"

func statusMessage(a int) {
	switch a {
	case 1:
		fmt.Println("Hello 1")
	case 2:
		fmt.Println("Hello 2")
	}
}