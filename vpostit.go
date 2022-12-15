package vpostit

import "context"

// Note represents an online post-it
type Note struct {
	Title string //title of your note
	Body  string //your sticky note
	Info
	ID string //a unique identifier for your note
}

type Info struct {
	MadeDay    string //date of the birthday of your note
	LastChange string //last change of your note
}

type Repository interface {
	Create(context.Context, *Note) error
	Update(context.Context, *Note) error
	FindByID(context.Context, NoteID) (_ Note, found bool, _ error)
	DeletByID(context.Context, NoteID) error
}
