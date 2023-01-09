package vpostit

import (
	"context"
	"fmt"
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
	if repo.MyNotes == nil {
		repo.MyNotes = make(map[string]*Note)
	}
	NewNote.ID = strconv.Itoa(rand.Int())

	repo.MyNotes[NewNote.ID] = NewNote //ezt a cimket beadom a hutomnek, es hozzaparositom az eppen letrehozni kivant jegyzetemet

	//return fmt.Errorf("ID: %v Note: %v couldn't create your PostIT", NewNote.ID, NewNote)
	return nil
}

func (repo *InMemoryNoteRepository) Update(ctx context.Context, update *Note) error {
	//en egy hutoszekreny vagyok, es van egy update metodusom, amivel engedem hogy frissitsek a fagyisdobozaim tartalmat.
	//ez a metodus megeszik egy kontextust, egy meglevo doboz tartalmat(oldNote), es az uj doboz komplett tartalmat(Note)
	//es ha baj van kikop egy hibauzenetet

	_, found, _ := repo.FindByID(ctx, update.ID)

	if !found {
		fmt.Println("couldn't update, unknown ID")

	} else {
		repo.MyNotes[update.ID] = update
	}

	return nil
}

func (repo *InMemoryNoteRepository) FindByID(ctx context.Context, ID string) (note Note, found bool, err error) {
	//en vagyok a huto, es fagyis dobozokat(jegyzeteket) tarolok.
	//van egy kereso metodusom, ami megeszik egy kontextust es a keresendo doboz cimkejet(ID).
	//cserebe kikopi a megtalalt dobozt(Note-ot), egy igent vagy egy nemet attol fuggoen, h megtalalta-e, es egy hibauzit ha vmi baj tortent.

	note, found = repo.MyNotes[ID]

	if !found {
		return Note{}, false, nil
	}

	return *note, true, nil

}

func (repo *InMemoryNoteRepository) DeleteByID(context.Context, string) error {

	return nil
}
