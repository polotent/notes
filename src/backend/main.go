package main

import (
	"log"
	"backend/router"
	"backend/controller"
	"backend/service"
	"backend/repository"
	"backend/db"
	"github.com/gin-gonic/gin"
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

	ginDispatcher := gin.New()
	router := router.NewGinRouter(ginDispatcher)
	router.GET("/api/notes", noteController.ReadAllNotes)
	// router.POST("api/notes", noteController.CreateNote)

	if err := router.Serve("8080"); err != nil {
		log.Fatalf("Server launch failed: %s", err)
	}
	// mux := http.NewServeMux()
	// router.Init(mux)
	// server := &http.Server{
	// 	Addr: "localhost:8080",
	// 	Handler: mux,
	// }

	// if err := server.ListenAndServe(); err != nil {
	// 	log.Fatalf("Server launch failed: %s", err)
	// }
}
