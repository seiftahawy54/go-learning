package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func (userNote *Note) PrintNoteContent() {
	fmt.Print("\nThis is your Note: \n", userNote.Title, "\n", userNote.Content, "\n", userNote.CreatedAt, "\n")
}

func (note Note) SaveNote() error {
	fileName := fmt.Sprint(note.Title, ' ', note.CreatedAt.Format("dd-mm-yyyy"))
	fileName = strings.ReplaceAll(fileName, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note) {
	newNote := Note{
		Title:   title,
		Content: content,
		CreatedAt: time.Now(),
	}

	return newNote
}
