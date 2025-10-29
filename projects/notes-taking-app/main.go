package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes-taking-app/notes"
)

func main() {

	// var notesData notes.Notes

	title := getUserInput("Enter the title")
	content := getUserInput("Please enter the content")

	notesData, err := notes.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	notesData.PrintOutPut()
}

func getUserInput(propmt string) string {

	fmt.Println(propmt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text

}
