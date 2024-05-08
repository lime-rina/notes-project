package collection

import (
	"fmt"

	"example.com/main/fileop"
	"example.com/main/note"
)

type Collection struct {
	notes []note.Note
}

func New() *Collection {
	return &Collection{
		notes: make([]note.Note, 0),
	}
}

func (c *Collection) AddNote() {
	newNote := note.AddNote()
	c.notes = append(c.notes, *newNote)
	fmt.Printf("Note %s added to collection\n", newNote.Id)
}

func (c *Collection) PrintNotes() {
	for _, n := range c.notes {
		n.Print()
	}
}

func (c *Collection) FindNoteById(id string) int {
	for idx, note := range c.notes {
		if note.Id == id {
			return idx
		}
	}
	return -1
}

func (c *Collection) isIndexValid() int {
	var id string
	fmt.Println("Enter a valid note ID: ")
	fmt.Scan(&id)
	idx := c.FindNoteById(id)
	return idx
}
func (c *Collection) DeleteNote() {
	idx := c.isIndexValid()
	if idx > -1 {
		c.notes = append(c.notes[:idx], c.notes[idx+1:]...)
		fmt.Print("Note deleted from collection\n")
	} else {
		fmt.Print("Note not found\n")
	}
}

func (c *Collection) EditNote() {
	idx := c.isIndexValid()
	if idx > -1 {
		c.notes[idx].EditNote()
		fmt.Print("Note deleted from collection\n")
	} else {
		fmt.Print("Note not found\n")
	}
}

// func (c *Collection) readNotesFromFile() {

// }

func (c *Collection) WriteNotesToFile() {
	file := ""
	fmt.Printf("Please enter a text file you want to write to: ")
	fmt.Scan(&file)
	for _, n := range c.notes {
		fileop.WriteIntoFile(file, &n)
	}
}
