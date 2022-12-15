package vpostit

import (
	"context"
	"time"
)

// Note represents an online post-it
type Note struct {
	Title string //title of your note
	Body  string //your sticky note
	Info
	ID string //a unique identifier for your note
}
type Time struct{}

type Info struct {
	MadeDay    time.Time //date of the birthday of your note
	LastChange time.Time //last change of your note
}

type Repository interface {
	Create(context.Context, *Note) error
	Update(context.Context, *Note) error
	FindByID(context.Context, NoteID) (_ Note, found bool, _ error)
	DeletByID(context.Context, NoteID) error
}

func (n Note) Create(context.Context, *Note) error {
	return nil
}

func (n Note) Update(context.Context, *Note) error {
	return nil
}
