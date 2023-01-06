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
	Create(context.Context, *Note, []MyNotes) error
	Update(context.Context, *Note) error
	FindByID(ctx context.Context, ID *Note) (_ Note, found bool, _ error)
	DeleteByID(ctx context.Context, ID string) error
}

type InMemoryNoteRepository struct {
	Service Repository
}

func (repo *InMemoryNoteRepository) Create(context.Context, *Note) error {

	NewNote := &Note{}
	ID := NewNote.ID

	var n = make(MyNotes) //I give the Note to a map, where the key is the ID and the value is the Note struct
	n[ID] = NewNote

	MyFluffyNotes = append(MyFluffyNotes, n) //here I give my map of the Note to the struct of all the Notes

	return nil
}

func (repo InMemoryNoteRepository) Update(ctx context.Context, oldNote, update *Note) error {

	*oldNote = *update
	return nil
}

func (repo *InMemoryNoteRepository) FindByID(ctx context.Context, ID *Note) (foundedNote Note, found bool, err error) {
	found = false
	for _, note := range MyFluffyNotes {
		noteID := note[ID]
		if noteID == ID { //invalid operation: cannot compare noteID == ID (mismatched types *Note and string)
			found = true
			break
		}
	}

	if !found {
		err = errors.New("oh fluff, couldn't find your note")
	}
	return foundedNote, found, err
}

/*
func (repo *InMemoryNoteRepository) DeleteByID(context.Context, string) error {
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
