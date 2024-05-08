package fileop

import (
	"bufio"
	"fmt"
	"os"

	"example.com/main/note"
)

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not open file")
		return nil
	}
	return file
}
func GetFileContent(path string) string {
	file := openFile(path)
	if file != nil {
		reader := bufio.NewReader(file)

		text, _ := reader.ReadString('\n')

		file.Close()
		return text
	}
	return ""
}

func WriteIntoFile(path string, note *note.Note) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	text := note.ConvertNoteToJSON()

	text += "\n"

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully wrote to file")
}
