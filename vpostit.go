package vpostit

import (
	"context"
	"time"
)

// Note represents an online post-it
type Note struct {
	Title string `json:"TITLE"` //title of your note
	Body  string `json:"BODY"`  //your sticky note
	Info  `json:"INFO"`
	ID    string `json:"ID"` //a unique identifier for your note
}

type Info struct {
	MadeDay    time.Time //date of the birthday of your note
	LastChange time.Time //last change of your note
}

type Repository interface {
	Create(context.Context, *Note) error
	Update(context.Context, *Note) error
	// FindByID(context.Context, NoteID) (_ Note, found bool, _ error)
	// DeletByID(context.Context, NoteID) error
}

type InMemoryNoteRepository struct {
	Service Repository
}

func (in InMemoryNoteRepository) Create(context.Context, *Note) error {
	return nil
}

func (in InMemoryNoteRepository) Update(context.Context, *Note) error {
	return nil
}
