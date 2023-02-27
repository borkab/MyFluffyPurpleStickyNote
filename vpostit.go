package vpostit

import (
	"context"
	"time"
)

// type MyNotes map[string]*Note
type NoteDTO struct {
	reqBody   string `json: "reqBody"`
	respBody  string `json: "respBody"`
	ID        string `json: "id"`
	fromQuery string `json: "fromQuery"`
}

// Note represents an online post-it
type Note struct { // TODO: erase json tags and make a NoteDTO instead.
	Title string //title of your note
	Body  string //your sticky note
	Info
	ID string //a unique identifier for your note
}

type Info struct {
	MadeDay  time.Time `json:"MadeDay"` //date of the birthday of your note
	UpdateAt time.Time `json:"Update"`  //last change of your note
}

// declaring the repository interface allows us to easily
// swap out the actual implementation, enforcing loose coupling.
type NoteRepository interface {
	Create(context.Context, *Note) error
	Update(context.Context, *Note) error
	FindByID(ctx context.Context, ID string) (_ Note, found bool, _ error)
	DeleteByID(ctx context.Context, ID string) error
	FindAllNow(ctx context.Context) ([]Note, error)
}
