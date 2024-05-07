package main

import (
	"example.com/main/collection"
	"example.com/main/menu"
)

func main() {
	collection := collection.New()
	menu.Menu(collection)
}
