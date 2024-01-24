package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)

}

func New(title, content string) (Note, error) {

	if content == "" || title == "" {
		return Note{}, errors.New("Invalid Input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}
