package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error

} 

type displayer interface {
	Display()

} 

type outputable interface {
	saver
	displayer
}

// type outputable interface {
// 	Save() error
// 	Display()

// }


func main() {

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return 
	}

	outputData(&todo)

	if err != nil {
		fmt.Println(err)
		return 
	}


	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return 
	}


	outputData(&userNote)

	printSomething(1)
	printSomething(1.5)
	printSomething("test")
	printSomething(todo)

}

func printSomething(value interface{}) {
	switch value.(type) {
	case int: 
		fmt.Println("Int")

	case string: 
		fmt.Println("string")

	default: 
		fmt.Println("other")
	}
	fmt.Println(value)
}


func outputData(data outputable) {
	data.Display()
	saveData(data)
}

func saveData(data saver) error{
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil

}


func getNoteData() (string,string) {
	title := getUserInput("Note Title: ")

	
	content := getUserInput("Note Content: ")


	return title, content
}

func getUserInput(prompt string) (string){

	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")


	return text

}
