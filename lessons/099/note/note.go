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
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Title: %s\n", note.Title)
	fmt.Printf("Content: %s\n", note.Content)
	//fmt.Printf("Created at: %s\n", note.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename)
	filename = fmt.Sprintf("%s.json", filename)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}
