package notes

import (
	"errors"
	"fmt"
	"time"
)

type Notes struct {
	title     string
	content   string
	createdAt time.Time
}

func (n Notes) PrintOutPut() {
	fmt.Printf("Title of the notes %v and description content %v", n.title, n.content)
}

func New(title, content string) (Notes, error) {

	if title == "" || content == "" {
		return Notes{}, errors.New("please fill the fields")
	}
	return Notes{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
