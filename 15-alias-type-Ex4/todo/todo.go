package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(item string) (Todo, error) {
	if item == "" {
		return Todo{}, errors.New("invalid Todo input")
	}
	return Todo{
		Text: item,
	}, nil
}
func (t Todo) Display() {
	fmt.Println(t.Text)
}
func (t Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(t)
	if err != nil {
		return errors.New("Todo save failed")
	}
	return os.WriteFile(fileName, json, 0644)
}
