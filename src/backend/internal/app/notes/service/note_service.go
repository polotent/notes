package service

import (
	"github.com/polotent/notes/src/backend/internal/app/notes/domain"
	"github.com/polotent/notes/src/backend/internal/app/notes/repository"
)

type NoteService interface {
	ReadAllNotes() ([]*domain.Note, error)
}

type noteService struct {
	noteRepository repository.NoteRepository
}

func (ns *noteService) ReadAllNotes() ([]*domain.Note, error) {
	return ns.noteRepository.ReadAllNotes()
}

func NewNoteService(nr repository.NoteRepository) NoteService {
	return &noteService{
		noteRepository: nr,
	}
}
