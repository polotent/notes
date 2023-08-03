package main

import (
	"log"
	"backend/router"
	"backend/controller"
	"backend/service"
	"backend/repository"
	"backend/db"
)

var (
	noteRepository = repository.NewNoteRepository()
	noteService = service.NewNoteService(noteRepository)
	noteController = controller.NewNoteController(noteService)
)

func main() {
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Fatalf("Database connection close failed: %s", err)
		}
	}()

	ginRouter := router.NewGinRouter()
	server := router.NewServer(ginRouter, noteController)
	server.Init()
	if err := server.Serve("8080"); err != nil {
		log.Fatalf("Server launch failed: %s", err)
	}
}
