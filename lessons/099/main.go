package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nikosmpi/note/note"
	"github.com/nikosmpi/note/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {
	printSomething(12)
	printSomething(12.23423)
	printSomething("Hello World")
	printSomething(true)
	title, content := getNoteData()
	todoText := getUserInput("Todo content:")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = outputData(todo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = outputData(userNote)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func printSomething(value interface{}) {
	if v, ok := value.(int); ok {
		fmt.Println("Integer:", v)
	} else if v, ok := value.(float64); ok {
		fmt.Println("Float:", v)
	} else if v, ok := value.(string); ok {
		fmt.Println("String:", v)
	} else if v, ok := value.(bool); ok {
		fmt.Println("Boolean:", v)
	} else {
		fmt.Println("Unknown type:", value)
	}
	//switch value.(type) {
	//case int:
	//	fmt.Println("Integer:", value)
	//case float64:
	//	fmt.Println("Float:", value)
	//case string:
	//	fmt.Println("String:", value)
	//case bool:
	//	fmt.Println("Boolean:", value)
	//default:
	//	fmt.Println("Unknown type:", value)
	//}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Note saved successfully!")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
