package main

import (
	"errors"
	"fmt"
	"unicode"
)

func allowOnlyCharacter(a string) error {
	for _, char := range a {
		if !unicode.IsLetter(char) {
			return errors.New("input contains non-alphabetic characters")
		}
	}
	return nil
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)

	// Call the validation function and handle the error
	err := allowOnlyCharacter(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid input:", input)
	}
}
