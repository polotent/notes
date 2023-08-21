package main

import (
	"log"

	notes "github.com/polotent/notes/src/backend/internal/app/notes"
)

func main() {
	notesApp, err := notes.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	err = notesApp.Run()
	if err != nil {
		log.Fatal(err)
	}
}
