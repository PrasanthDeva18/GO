package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Notes struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

// In go json marshal package directly converts the struct data to json data but make sure struct data members should be publically accessible

func (n Notes) PrintOutPut() {
	fmt.Printf("Title of the notes %v and description content %v", n.Title, n.Content)
}

// struct tags are the metadata
// type struct -> title string `json:"name value -> you want"`

func (n Notes) SaveJsonData() error {
	filename := strings.ReplaceAll(n.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}

func New(title, content string) (Notes, error) {

	if title == "" || content == "" {
		return Notes{}, errors.New("please fill the fields")
	}
	return Notes{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
