package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)


func main() {
	// fmt.Println("Welcome to the greatest notes app")
	// title, content, err := insertNotes()
	// createdNote := note.New(title, content)

	// if err != nil {
	// 	fmt.Print(err)
	// }

	fmt.Println(time.Now().Date())

	// createdNote.PrintNoteContent()
	// err = createdNote.SaveNote()
	// fmt.Print(err)
}

func printAnyThing(data ... any) {
	fmt.Print(data)
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')

	if err != nil || userInput == "" {
		return "", errors.New("Invalid input")
	}

	// Remove auto-added line breaks
	userInput = strings.TrimSuffix(userInput, "\n")
	// Remove windows line break
	userInput = strings.TrimSuffix(userInput, "\r")

	return userInput, nil
}


func insertNotes() (string, string, error) {
	title, err := getUserInput("Please enter your note title: ")

	if err != nil {
		fmt.Println()
		return "", "", errors.New("Please enter a valid note title") 
	}

	content, err := getUserInput("Please enter the notes content: ")

	if err != nil {
		return "", "", errors.New("Please enter a valid note content")
	}

	return title, content, nil
}

