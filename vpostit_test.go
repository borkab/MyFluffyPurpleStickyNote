package vpostit

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
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
	}

	Note2 := &Note{
		Title: "Evenings TODO",
		Body:  "pick up toys, pick up clothes, set dishwasher, take out trash",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}
	/*
		Note3 := &Note{
			Title: "Some things",
			Body:  "dance, sing, eat",
			Info: Info{
				MadeDay:  time.Now(),
				UpdateAt: time.Now(),
			},
		}
	*/
	repo := InMemoryNoteRepository{}
	ctx := context.Background()

	repo.Create(ctx, Note1)
	repo.Create(ctx, Note2)
	//repo.Create(ctx, Note3)
	//fmt.Println(Note1.ID)
	//fmt.Println(Note2.ID)
	//fmt.Println(Note3.ID)

	//ID1 := Note1.ID
	//ID2 := Note2.ID
	ID3 := strconv.Itoa(rand.Int())

	repo.FindByID(ctx, Note1.ID)
	repo.FindByID(ctx, Note2.ID)
	repo.FindByID(ctx, ID3)

	UpdateNote1 := &Note{
		Title: "Morning Rituals",
		Body:  "brush teeth, wash face, drink water",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	UpdateNote2 := &Note{
		Title: "Evening Rituals",
		Body:  "brush teeth, have a shower, go to sleep",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	UpdateNote3 := &Note{
		Title: "Evening Rituals",
		Body:  "brush teeth, have a shower, go to sleep",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	repo.Update(ctx, UpdateNote1)
	repo.Update(ctx, UpdateNote2)
	repo.Update(ctx, UpdateNote3)
}

func TestCreate(t *testing.T) {
	newNote := &Note{
		Title: "Shopping list",
		Body:  "milk, coffee, pretzels",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
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
			MadeDay: time.Now(),
		},
	}

	update := &Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: Info{
			UpdateAt: time.Now(),
		},
	}
	repo := InMemoryNoteRepository{}
	ctx := context.Background()
	repo.Create(ctx, oldNote)
	got := repo.Update(ctx, update)
	//	too many arguments in call to repo.Update
	//	have (context.Context, *Note, *Note)
	//	want (context.Context, *Note)compiler

	if got != nil {
		t.Fatal("couldn't update note")
	}
}

func TestFoundByID(t *testing.T) {

	Note1 := &Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note2 := &Note{
		Title: "My Today's TODO list",
		Body:  "go shopping, do housework, learn GO",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note3 := &Note{
		Title: "My Today's Housework list",
		Body:  "make laundry, cook lunch, pick up toys, hoover everywhere",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	repo := InMemoryNoteRepository{}
	ctx := context.Background()
	repo.Create(ctx, Note1)
	repo.Create(ctx, Note2)
	repo.Create(ctx, Note3)
	//fmt.Println(Note1.ID, Note2.ID, Note3.ID)

	t.Run("if found", func(t *testing.T) {
		ID := Note2.ID

		_, found, err := repo.FindByID(ctx, ID)

		if !found {
			t.Fatal("couldn't find this note")
		}

		if err != nil {
			t.Fatal("there is a houge problem")
		}
	})

	t.Run("if not found", func(t *testing.T) {
		ID := strconv.Itoa(rand.Int())

		_, found, err := repo.FindByID(ctx, ID)

		if !found {
			fmt.Println("test for unknown ID: !found: couldn't find this note")
		}

		if err != nil {
			fmt.Println(errors.New("test for unknown ID: err: your note could not be found"))
		}
	})
}
