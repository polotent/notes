package repository

import (
	"log"
	"backend/domain"
	"database/sql"
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
		log.Fatal(err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var note domain.Note
	for rows.Next() {
		if err := rows.Scan(&note.Id, &note.Title, &note.Description); err != nil {
			log.Fatal(err)
		}
		notes = append(notes, &note)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return notes, nil
}

func NewNoteRepository(db *sql.DB) NoteRepository {
	return &noteRepository{
		db: db,
	}
}
