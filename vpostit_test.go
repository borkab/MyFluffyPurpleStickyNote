package vpostit

import (
	"context"
	"testing"
	"time"
)

func TestInMemoryNoteRepository_smoke(t *testing.T) {

	Note1 := &Note{
		Title: "Mornings TODO",
		Body:  "make laundry, cook lunch, clean dining table, wash dishes",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
		ID: "",
	}

	Note2 := &Note{
		Title: "Evenings TODO",
		Body:  "pick up toys, pick up clothes, set dishwasher, take out trash",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
		ID: "",
	}

	repo := InMemoryNoteRepository{}
	ctx := context.Background()

	repo.Create(ctx, Note1)
	repo.Create(ctx, Note2)

	ID1 := Note1.ID
	ID2 := Note2.ID
	repo.FindByID(ctx, ID1)
	repo.FindByID(ctx, ID2)
}

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

	repo := InMemoryNoteRepository{}
	ctx := context.Background()

	got := repo.Create(ctx, newNote)

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
	MyFluffyNotes := []*Note{
		&Note{
			Title: "My Today's Shopping list",
			Body:  "milk, coffee, bagels and more",
			ID:    "fluff.0001",
		},
		&Note{
			Title: "My Today's TODO list",
			Body:  "go shopping, do housework, learn GO",
			ID:    "fluff.0002",
		},
		&Note{
			Title: "My Today's Housework list",
			Body:  "make laundry, cook lunch, pick up toys, hoover everywhere",
			ID:    "fluff.0003",
		},
	}

	ID := "fluff.0002"

	memo := InMemoryNoteRepository{}
	ctx := context.Background()

	_, found, err := memo.FindByID(ctx, ID, MyFluffyNotes)

	if found != true {
		t.Fatal("couldn't find this note")
	}

	if err != nil {
		t.Fatal("")
	}
}
*/
