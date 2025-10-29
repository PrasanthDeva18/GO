package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// var a, b string

	// fmt.Scan(&a)
	// fmt.Scan(&b)

	reader := bufio.NewReader(os.Stdin)

	//rune is called in go
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	text = strings.TrimSuffix(text, "\n")

	text = strings.TrimSuffix(text, "\r")

	fmt.Println(text)

}
