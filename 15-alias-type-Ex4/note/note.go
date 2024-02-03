package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid Note input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
func (n Note) Display() {
	fmt.Printf("\nTitle:\n%s\nContent:\n%s\n", n.Title, n.Content)
}
func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "-")
	fileName = strings.ReplaceAll(fileName, ",", "-")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)
	if err != nil {
		return errors.New("Note save failed")
	}
	return os.WriteFile(fileName, json, 0644)
}
