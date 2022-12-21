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
			MadeDay:    time.Now(),
			LastChange: time.Now(),
		},
		ID: "fluff.0001",
	}
	memo := InMemoryNoteRepository{}
	ctx := context.Background()
	got := memo.Create(ctx, newNote)

	if got != nil {
		t.Fatal("couldn't create new note")
	}
}

func TestUpdate(t *testing.T) {
	updateNote := &Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: Info{
			LastChange: time.Now(),
		},
	}
	memo := InMemoryNoteRepository{}
	ctx := context.Background()
	got := memo.Update(ctx, updateNote)

	if got != nil {
		t.Fatal("couldn't create new note")
	}
}
