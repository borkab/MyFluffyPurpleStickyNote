package vpostit

import (
	"context"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	newNote := &Note{
		Title: "Shopping list",
		Body:  "milk, coffee, pretzels",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
		ID: "fluff.0001",
	}
	MyFluffyNotes := []*Note{}

	memo := InMemoryNoteRepository{}
	ctx := context.Background()
	got := memo.Create(ctx, newNote, MyFluffyNotes)

	if got != nil {
		t.Fatal("couldn't create new note")
	}
}

func TestUpdate(t *testing.T) {
	oldNote := &Note{
		Title: "Shopping list",
		Body:  "milk, coffee, pretzels",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
		ID: "fluff.0001",
	}

	update := Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: Info{
			UpdateAt: time.Now(),
		},
	}

	ctx := context.Background()
	got := oldNote.Update(ctx, update)

	if got != nil {
		t.Fatal("couldn't update note")
	}
}

/*
func TestFoundByID(t *testing.T) {
	note := &Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: Info{
			UpdateAt: time.Now(),
		},
		ID: "fluff.0001",
	}
	var ID *Note
	ID = note.ID
	//how should I make frome string a *Note type,
	// that it accepts as a parameter?

	memo := InMemoryNoteRepository{}
	ctx := context.Background()

	_, found, err := memo.FindByID(ctx, ID)

	if found != true {
		t.Fatal("couldn't find this note")
	}

	if err != nil {
		t.Fatal("")
	}
}
*/
