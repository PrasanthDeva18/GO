package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string
}

func (t Todo) GetUserOp() {
	fmt.Println("Todo Content :%v", t.Content)
}

func (t Todo) SaveData() error {
	filename := "todo-list" + ".json"

	json, err := json.Marshal(t)

	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)

}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("Fields are missing")
	}
	return Todo{
		Content: content,
	}, nil
}
