package vpostit

import (
	"context"
	"errors"
	"time"
)

// MyFluffyNotes holds all of my Notes in a slice of type MyNotes
var MyFluffyNotes []MyNotes

// MyNotes is a map of a *Note, for easier manipulating as search and delete by ID(key)
type MyNotes map[string]*Note

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

func (in *InMemoryNoteRepository) Create(context.Context, *Note, map[Note]*Note) error {
	//context.TODO()

	NewNote := &Note{}
	ID := NewNote.ID
	MyNotes[ID] = *NewNote //invalid operation: MyNotes[Note.ID]
	//(MyNotes is not a generic type)
	N := MyNotes[ID] * NewNote

	MyFluffyNotes = append(MyFluffyNotes, MyNotes)

	return nil
}

func (oldNote *Note) Update(ctx context.Context, update Note) error {
	context.TODO()
	*oldNote = update
	return nil
}

func (in *InMemoryNoteRepository) FindByID(context.Context, string, []*Note) (ID string, found bool, err error) {
	found = false
	for _, note := range MyFluffyNotes {
		if note.ID == ID {
			found = true
			break
		}
	}

	if !found {
		err = errors.New("oh fluff, couldn't find your note")
	}
	return ID, found, err
}

/*
func (in *InMemoryNoteRepository) DeleteByID(context.Context, string) error {
	//your code
	mynote := &Note{}
	ctx := context.TODO()

	ID := mynote.ID
	var err error
	_, found, _ := in.FindByID(ctx, ID, MyFluffyNotes)
	if !found {
		err = errors.New("oh fluff, couldn't find your note")
	} else {
		for i := range MyFluffyNotes {
			if MyFluffyNotes[i{ID}] == mynote.ID {
				return i
			}
		}
	}

	return err
}
*/
