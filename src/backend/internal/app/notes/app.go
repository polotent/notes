package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/polotent/notes/src/backend/internal/app/notes/config"
	"github.com/polotent/notes/src/backend/internal/app/notes/controller"
	"github.com/polotent/notes/src/backend/internal/app/notes/database"
	"github.com/polotent/notes/src/backend/internal/app/notes/repository"
	"github.com/polotent/notes/src/backend/internal/app/notes/router"
	"github.com/polotent/notes/src/backend/internal/app/notes/server"
	"github.com/polotent/notes/src/backend/internal/app/notes/service"
)

type App struct {
	server server.Server
	db     *sql.DB
}

func (app *App) Run() error {
	defer app.close()
	if err := app.server.Run("8080"); err != nil {
		log.Printf("Server launch failed: %s\n", err)
		return err
	}

	return nil
}

func (app *App) close() {
	if err := app.db.Close(); err != nil {
		log.Fatalf("Database connection closing failed: %s", err)
	}
	log.Fatalf("Close connection to database")
}

func NewApp() (*App, error) {
	notesApp := &App{}

	appConfig, _ := config.LoadAppConfig()

	db, err := database.ConnectDatabase(appConfig)
	if err != nil {
		log.Printf("Failed to connect to database: %s\n", err)
		os.Exit(2)
	}
	log.Printf("Successfully connected to database with host=%s, port=%s\n", appConfig.Database.Host, appConfig.Database.Port)
	notesApp.db = db

	noteRepository := repository.NewNoteRepository(notesApp.db)
	noteService := service.NewNoteService(noteRepository)
	noteController := controller.NewNoteController(noteService)

	ginRouter := router.NewGinRouter()

	notesApp.server = server.NewServer(ginRouter, noteController)
	notesApp.server.Setup()

	return notesApp, nil
}
