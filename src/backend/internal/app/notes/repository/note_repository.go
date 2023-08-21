package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/polotent/notes/src/backend/internal/app/notes/domain"
)

type NoteRepository interface {
	ReadAllNotes() ([]*domain.Note, error)
}

type noteRepository struct {
	db *sql.DB
}

func (nr *noteRepository) ReadAllNotes() ([]*domain.Note, error) {
	var notes []*domain.Note
	rows, err := nr.db.Query("SELECT * FROM notes")
	if err != nil {
		log.Println(err)
		return nil, errors.New("Error executing query")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	var note domain.Note
	for rows.Next() {
		if err := rows.Scan(&note.Id, &note.Title, &note.Description); err != nil {
			log.Println(err)
			return nil, errors.New("Error reading table rows")
		}
		notes = append(notes, &note)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, errors.New("Error reading table rows")
	}

	return notes, nil
}

func NewNoteRepository(db *sql.DB) NoteRepository {
	return &noteRepository{
		db: db,
	}
}
