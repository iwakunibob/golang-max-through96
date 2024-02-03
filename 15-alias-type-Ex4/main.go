package main

import (
	"bufio"
	"fmt"
	"max/note/note"
	"max/note/todo"
	"os"
	"strings"
)

type saver interface { // An interface is a contract
	Save() error // All types using it must have a Save method
} // Note error/nil is returned

type outputable interface { // Embedded interface can contain
	saver     // other multiple interfaces
	Display() // other multiple methods
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	todoText := getTodoData()
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Saving note succeeded")
	err = outputData(userTodo)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data) // Embed interface func
}
func saveData(data saver) error { // Interface func
	err := data.Save()
	if err != nil {
		return err
	}
	return nil
}

// The Any value is allowed type interface{} "any" is an alias
func printAnything(value any) {
	fmt.Println(value)
}

func getTodoData() string {
	return getUserInput("Todo text: ")
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}
func getUserInput(prompt string) string {
	fmt.Printf("%s ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	// fmt.Scanln(&value) // Can't use because sentence not word
	return text
}
