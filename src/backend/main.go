package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/server"
	"backend/service"
	"log"
)

var (
	noteRepository = repository.NewNoteRepository()
	noteService    = service.NewNoteService(noteRepository)
	noteController = controller.NewNoteController(noteService)
)

func main() {
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Fatalf("Database connection close failed: %s", err)
		}
	}()

	ginRouter := router.NewGinRouter()
	server := server.NewServer(ginRouter, noteController)
	server.Setup()
	if err := server.Run("8080"); err != nil {
		log.Fatalf("Server launch failed: %s", err)
	}
}
