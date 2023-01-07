package vpostit

import (
	"context"
	"math/rand"
	"strconv"
	"time"
)

//type MyNotes map[string]*Note

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
	// MyNotes is a map of a *Note, for easier manipulating as search and delete by ID(key)
	MyNotes map[string]*Note
}

func (repo *InMemoryNoteRepository) Create(ctx context.Context, NewNote *Note) error {
	repo.MyNotes = make(map[string]*Note)
	NewNote.ID = strconv.Itoa(rand.Int())

	repo.MyNotes[NewNote.ID] = NewNote //ezt a cimket beadom a hutomnek, es hozzaparositom az eppen letrehozni kivant jegyzetemet

	//return fmt.Errorf("ID: %v Note: %v couldn't create your PostIT", NewNote.ID, NewNote)
	return nil
}

/*
func (repo InMemoryNoteRepository) Update(ctx context.Context, oldNote, update *Note) error {

	*oldNote = *update
	return nil
}
*/

/*
func (repo *InMemoryNoteRepository) FindByID(ctx context.Context, ID string) (foundedNote Note, found bool, err error) {
	found = false
	for _, note := range repo.MyNotes { //vegigmegyek repo.MyNotes map osszes elemen, es amikor az elsohoz ertem,

		noteID := note.ID
		if noteID == ID { //megnezem h a kez megegyezik-e az altalam keresett ID-val. ha igen,
			found = true //megtalaltam
			break        //ezert nem keresek tovabb
		}
	}

	if !found {
		err = errors.New("oh fluff, couldn't find your note")
	}
	return foundedNote, found, err
}
*/

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
