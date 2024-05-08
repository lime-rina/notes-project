package note

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/teris-io/shortid"
)

type Note struct {
	title       string
	description string
	Id          string
}

func New(t, d string) (*Note, error) {
	if t == "" || d == "" {
		return nil, errors.New("invalid note data")
	}
	id, err := shortid.Generate()
	if err != nil {
		return nil, err
	}
	return &Note{
		Id:          id,
		title:       t,
		description: d,
	}, nil
}

func (n *Note) Print() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("ID: %s\nTitle: %s\nDescription: %s\n", n.Id, n.title, n.description)
}

func InputNoteData(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func (n *Note) EditNote() {

	fmt.Print("What do you want to edit - the title or the description (t/d) ")
	var answer string
	fmt.Scan(&answer)
	switch answer {
	case "t":
		n.title = InputNoteData("Title: ")
	case "d":
		n.description = InputNoteData("Description: ")
	default:
		fmt.Println("Unknown answer- you can only choose between title and description")
	}
}

func (n *Note) DeleteNote() {
	n.title = ""
	n.description = ""
	n.Id = ""
}

func AddNote() *Note {
	title := InputNoteData("Title: ")
	description := InputNoteData("Description: ")

	newNote, _ := New(title, description)
	return newNote
}

func (n *Note) ConvertNoteToJSON() string {
	return fmt.Sprintf("{title: %s , description: %s, id: %s}\n", n.title, n.description, n.Id)
}
