package controller

import (
	"encoding/json"
	"net/http"

	"github.com/polotent/notes/src/backend/internal/app/notes/errors"
	"github.com/polotent/notes/src/backend/internal/app/notes/service"
)

type NoteController interface {
	ReadAllNotes(w http.ResponseWriter, r *http.Request)
}

type noteController struct {
	noteService service.NoteService
}

func (nc *noteController) ReadAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := nc.noteService.ReadAllNotes()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{ErrorMessage: "Error getting notes"})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ControllerError{ErrorMessage: "Error marshalling notes"})
		return
	}
}

func NewNoteController(ns service.NoteService) NoteController {
	return &noteController{
		noteService: ns,
	}
}
