package controller

import (
	"backend/domain"
	"backend/service"
)

type NoteController interface {
	ReadAllNotes() ([]*domain.Note, error)
}

type noteController struct {
	noteService service.NoteService
}

func (nc *noteController) ReadAllNotes() ([]*domain.Note, error) {
	return nc.noteService.ReadAllNotes()
}

func NewNoteController(ns service.NoteService) NoteController {
	return &noteController{
		noteService: ns,
	}
}