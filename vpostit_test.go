package vpostit

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/adamluzsi/testcase/assert"
)

func TestInMemoryNoteRepository_smoke(t *testing.T) {
	n1 := Note{
		Title: "Mornings TODO",
		Body:  "make laundry, cook lunch, clean dining table, wash dishes",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	n2 := Note{
		Title: "Evenings TODO",
		Body:  "pick up toys, pick up clothes, set dishwasher, take out trash",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	repo := InMemoryNoteRepository{}
	ctx := context.Background()

	t.Log("Create")
	assert.NoError(t, repo.Create(ctx, &n1))
	assert.NotEmpty(t, n1.ID)

	assert.NoError(t, repo.Create(ctx, &n2))
	assert.NotEmpty(t, n2.ID)

	t.Log("FindByID - n1")
	gotNote, found, err := repo.FindByID(ctx, n1.ID)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, n1, gotNote)

	t.Log("FindByID - n2")
	gotNote, found, err = repo.FindByID(ctx, n2.ID)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, n2, gotNote)

	t.Log("Update")
	n1.Title = "foo/bar/baz"
	assert.NoError(t, repo.Update(ctx, &n1), "updating note1 should be possible")
	gotNote, found, err = repo.FindByID(ctx, n1.ID)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, n1, gotNote)

	t.Log("DeleteByID")
	assert.NoError(t, repo.DeleteByID(ctx, n1.ID), "deleting note1 should be possible")
	gotNote, found, err = repo.FindByID(ctx, n1.ID)
	assert.NoError(t, err)
	assert.False(t, found, "note1 should be deleted")
	assert.Empty(t, gotNote)

	gotNote, found, err = repo.FindByID(ctx, n2.ID)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, n2, gotNote, "note2 should be the same regardless of the deletion of note1")

	assert.Error(t, repo.Update(ctx, &n1), "updating a deleted entity should yield an error")
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

	repo := InMemoryNoteRepository{}
	ctx := context.Background()
	repo.Create(ctx, oldNote) //letrehozok egy oldNote nevu jegyzetet a repoban

	update := &Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: Info{
			UpdateAt: time.Now(),
		},
		ID: oldNote.ID,
	}

	update2 := &Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: Info{
			UpdateAt: time.Now(),
		},
		ID: strconv.Itoa(rand.Int()),
	}

	t.Run("if found", func(t *testing.T) {
		got := repo.Update(ctx, update)

		if got != nil {
			t.Fatal("couldn't update note")
		}
	})

	t.Run("if not found", func(t *testing.T) {
		got := repo.Update(ctx, update2)

		if got == nil {
			t.Fatal("how can you update an unknown note?")
		}
	})
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
			fmt.Printf("!found: couldn't find this note")
		}

		if err != nil {
			fmt.Printf("err: your note could not be found")
		}
	})
}

func TestDeleteByID(t *testing.T) {

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

	Note4 := &Note{
		Title: "Learning TODO",
		Body:  "maps, interfaces, RDD, TDD",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
		ID: strconv.Itoa(rand.Int()),
	}

	repo := InMemoryNoteRepository{}
	ctx := context.Background()
	repo.Create(ctx, Note1)
	repo.Create(ctx, Note2)
	repo.Create(ctx, Note3)

	repo.DeleteByID(ctx, Note1.ID)
	repo.DeleteByID(ctx, Note2.ID)
	repo.DeleteByID(ctx, Note3.ID)
	repo.DeleteByID(ctx, Note4.ID)
}
