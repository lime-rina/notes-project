package menu

import (
	"fmt"

	"example.com/main/collection"
)

func displayMenu(userChoice *uint) {
	fmt.Print("-------------------------------\n")
	fmt.Println("1. Add a new note")
	fmt.Println("2. Edit a note (id)")
	fmt.Println("3. Delete a note (id)")
	fmt.Println("4. Print notes")
	fmt.Println("5. Exit")
	fmt.Print("-------------------------------\n")
	fmt.Scan(userChoice)
}

func Menu(collection *collection.Collection) {
	var userInput uint = 0
	for {
		displayMenu(&userInput)
		switch userInput {
		case 1:
			collection.AddNote()
		case 2:
			collection.EditNote()
		case 3:
			collection.DeleteNote()
		case 4:
			collection.PrintNotes()
		case 5:
			fmt.Println("You chose to exit")
			return
		default:
			fmt.Println("Please enter a valid option")
		}

	}
}
