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
	Desc      string
	CreatedAt time.Time
}

func (n Notes) GetUserOp() {
	fmt.Println("Note Title %v and Description of Notes: %v", n.Title, n.Desc)
}

func (n Notes) SaveData() error {
	filename := strings.ReplaceAll(n.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)

}

func New(title, desc string) (Notes, error) {

	if title == "" || desc == "" {
		return Notes{}, errors.New("Fields are missing")
	}
	return Notes{
		Title:     title,
		Desc:      desc,
		CreatedAt: time.Now(),
	}, nil
}
