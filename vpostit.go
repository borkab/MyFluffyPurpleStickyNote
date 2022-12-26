package vpostit

import (
	"context"
	"errors"
	"fmt"
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
	MadeDay  time.Time //date of the birthday of your note
	UpdateAt time.Time //last change of your note
}

// declaring the repository interface allows us to easily
// swap out the actual implementation, enforcing loose coupling.
type Repository interface {
	Create(context.Context, *Note) error
	Update(context.Context, *Note) error
	FindByID(ctx context.Context, ID *Note) (_ Note, found bool, _ error)
	DeleteByID(ctx context.Context, ID string) error
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

func (in *InMemoryNoteRepository) FindByID(ctx context.Context, ID *Note) (_ Note, found bool, _ error) {
	//what comes here?
	if !found {
		fmt.Println(errors.New("oh fluff, couldn't find your note"))
	}
	return
}

func (in *InMemoryNoteRepository) DeleteByID(context.Context, string) error {
	//your code
	mynote := &Note{}
	ctx := context.Background()

	ID := mynote.ID
	in.FindByID(ctx, ID) //how the hell could I convert string to *Note??
	//cannot use ID (variable of type string) as *Note value in argument to in.FindByID___compiler
	return nil
}
