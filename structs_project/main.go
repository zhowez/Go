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


func main() {

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return 
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return 
	}

	fmt.Println("Saving the note succeeded!")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return 
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return 
	}

	fmt.Println("Saving the note succeeded!")

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
