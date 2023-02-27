package memory

import (
	"context"
	"errors"
	"math/rand"
	"sort"
	"strconv"
	vpostit "vPOST-it"
)

type NoteRepository struct {
	// notes is a map of a *Note, for easier manipulating as search and delete by ID(key)
	notes map[string]*vpostit.Note
}

func (repo *NoteRepository) FindAllNow(ctx context.Context) ([]vpostit.Note, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	var notes []vpostit.Note
	for _, n := range repo.notes {
		notes = append(notes, *n)
	}
	sort.Slice(notes, func(i, j int) bool { //sorba rendezi a note-okat az ID-juk alapjan?
		return notes[i].ID < notes[j].ID
	})
	return notes, nil
}

func (repo *NoteRepository) Create(ctx context.Context, NewNote *vpostit.Note) error {
	if repo.notes == nil {
		repo.notes = make(map[string]*vpostit.Note)
	}
	NewNote.ID = strconv.Itoa(rand.Int())

	repo.notes[NewNote.ID] = NewNote //ezt a cimket beadom a hutomnek, es hozzaparositom az eppen letrehozni kivant jegyzetemet

	//return fmt.Errorf("ID: %v Note: %v couldn't create your PostIT", NewNote.ID, NewNote)
	return nil
}

func (repo *NoteRepository) Update(ctx context.Context, update *vpostit.Note) error {
	//en egy hutoszekreny vagyok, es van egy update metodusom, amivel engedem hogy frissitsek a fagyisdobozaim tartalmat.
	//ez a metodus megeszik egy kontextust, egy meglevo doboz tartalmat(oldNote), es az uj doboz komplett tartalmat(Note)
	//es ha baj van kikop egy hibauzenetet

	_, found, err := repo.FindByID(ctx, update.ID) //megnezem, h a frissiteni kivant fagyisdoboz benne van e pocakomban

	if !found { //ha nincs benne, kiirjuk hogy ismeretlen cimke
		err = errors.New("couldn't update, unknown ID")
	}

	repo.notes[update.ID] = update

	return err //visszadobunk egy nulla erteku hibauzit
}

func (repo *NoteRepository) FindByID(ctx context.Context, ID string) (vpostit.Note, bool, error) {
	//en vagyok a huto, es fagyis dobozokat(jegyzeteket) tarolok.
	//van egy kereso metodusom, ami megeszik egy kontextust es a keresendo doboz cimkejet(ID).
	//cserebe kikopi a megtalalt dobozt(Note-ot), egy igent vagy egy nemet attol fuggoen, h megtalalta-e, es egy hibauzit ha vmi baj tortent.
	note, found := repo.notes[ID]
	if !found {
		return vpostit.Note{}, false, nil
	}
	return *note, true, nil
}

func (repo *NoteRepository) DeleteByID(ctx context.Context, ID string) error {
	//egy tarolo vagyok, es jegyzeteket tarolnak bennem
	//van egy delete metodusom, ami beker egy kontextust es egy egyedi kodot,
	//majd visszaad egy hibauzenetet, ha valami balul sult el a torles soran

	_, found, err := repo.FindByID(ctx, ID) //megnezem h a torolni kivant jegyzet egyedi kodja bennem van e

	if !found {
		return errors.New("unknown ID")
	}
	delete(repo.notes, ID)
	return err
}
