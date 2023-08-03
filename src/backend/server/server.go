package server

import (
	"backend/controller"
	"backend/router"
)

type Server interface {
	Setup()
	Run(port string) error
}

type server struct {
	router         router.Router
	noteController controller.NoteController
}

func (server *server) Setup() {
	server.setNoteRoutes()
}

func (server *server) Run(port string) error {
	return server.router.Serve(port)
}

func (server *server) setNoteRoutes() {
	server.router.GET("/api/notes", server.noteController.ReadAllNotes)
}

func NewServer(router router.Router, noteController controller.NoteController) Server {
	return &server{
		router:         router,
		noteController: noteController,
	}
}
