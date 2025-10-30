package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/notes"
	"example.com/notes/todo"
)

// type

type saver interface {
	SaveData() error
}

func main() {
	fmt.Println("Developing notes taking app")

	title := getUserInp("Enter the title of notes")
	desc := getUserInp("Enter the desc of notes")
	tcontent := getUserInp("Enter the todo content")

	notesData, err := notes.New(title, desc)

	if err != nil {
		fmt.Println(err)
		return
	}

	notesData.GetUserOp()
	// err = notesData.SaveData()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	saveData(notesData)

	todoData, err := todo.New(tcontent)

	if err != nil {
		fmt.Println(err)
		return
	}

	todoData.GetUserOp()
	saveData(todoData)
	// err = todoData.SaveData()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

func saveData(s saver) {
	err := s.SaveData()

	if err != nil {
		fmt.Println("error in saving data")
	}

	fmt.Println("Saved data successfully")
}

func getUserInp(prompttext string) string {

	fmt.Println(prompttext)

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
