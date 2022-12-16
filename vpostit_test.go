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
	ctx := context.Background()
	got := Create(ctx, newNote)

	if got != nil {
		t.Fatal("couldn't create new note")
	}
}
