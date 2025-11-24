package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"content"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("Text: %s\n", todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}
