package server

import (
	"backend/controller"
	"backend/router"
)

type Server interface {
	Init()
	Serve(port string) error
}

type server struct {
	router         router.Router
	noteController controller.NoteController
}

func (appRouter *server) Init() {
	appRouter.setNoteRoutes()
}

func (appRouter *server) Serve(port string) error {
	return appRouter.router.Serve(port)
}

func (appRouter *server) setNoteRoutes() {
	appRouter.router.GET("/api/notes", appRouter.noteController.ReadAllNotes)
}

func NewServer(router router.Router, noteController controller.NoteController) Server {
	return &server{
		router:         router,
		noteController: noteController,
	}
}
