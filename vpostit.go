package vpostit

import (
	"context"
	"errors"
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

	_, found, _ := repo.FindByID(ctx, update.ID) //megnezem, h a frissiteni kivant fagyisdoboz benne van e pocakomban

	if !found { //ha nincs benne, kiirjuk hogy ismeretlen cimke
		fmt.Println("couldn't update, unknown ID")

	} else { //amugy meg frissitjuk ha megvan
		repo.MyNotes[update.ID] = update
	}

	return nil //visszadobunk egy nulla erteku hibauzit
}

func (repo *InMemoryNoteRepository) FindByID(ctx context.Context, ID string) (Note, bool, error) {
	//en vagyok a huto, es fagyis dobozokat(jegyzeteket) tarolok.
	//van egy kereso metodusom, ami megeszik egy kontextust es a keresendo doboz cimkejet(ID).
	//cserebe kikopi a megtalalt dobozt(Note-ot), egy igent vagy egy nemet attol fuggoen, h megtalalta-e, es egy hibauzit ha vmi baj tortent.
	note, found := repo.MyNotes[ID]
	if !found {
		return Note{}, false, nil
	}
	return *note, true, nil
}

func (repo *InMemoryNoteRepository) DeleteByID(ctx context.Context, ID string) error {
	//egy tarolo vagyok, es jegyzeteket tarolnak bennem
	//van egy delete metodusom, ami beker egy kontextust es egy egyedi kodot,
	//majd visszaad egy hibauzenetet, ha valami balul sult el a torles soran

	_, found, err := repo.FindByID(ctx, ID) //megnezem h a torolni kivant jegyzet egyedi kodja bennem van e

	if !found {
		return errors.New("unknown ID")
	}
	delete(repo.MyNotes, ID)
	return err
}
