package memory_test

import (
	"context"
	"fmt"
	"github.com/adamluzsi/testcase/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
	vpostit "vPOST-it"
	"vPOST-it/adapters/memory"
)

// _ valtozo tipusa NoteRepository ami egy interface a vpostit packageben.
// az erteke pedig egy a memory packageben talalhato NoteRepository interfacere mutato pointer
var _ vpostit.NoteRepository = &memory.NoteRepository{}

func TestInMemoryNoteRepository_smoke(t *testing.T) {
	n1 := vpostit.Note{
		Title: "Mornings TODO",
		Body:  "make laundry, cook lunch, clean dining table, wash dishes",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	n2 := vpostit.Note{
		Title: "Evenings TODO",
		Body:  "pick up toys, pick up clothes, set dishwasher, take out trash",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	repo := memory.NoteRepository{}
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
	newNote := &vpostit.Note{
		Title: "Shopping list",
		Body:  "milk, coffee, pretzels",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	repo := memory.NoteRepository{}
	ctx := context.Background()

	got := repo.Create(ctx, newNote)

	if got != nil {
		t.Fatal("couldn't create new note")
	}
}

func TestUpdate(t *testing.T) {
	note := &vpostit.Note{
		Title: "Shopping list",
		Body:  "milk, coffee, pretzels",
		Info: vpostit.Info{
			MadeDay: time.Now(),
		},
	}

	repo := memory.NoteRepository{}
	ctx := context.Background()
	repo.Create(ctx, note) //letrehozok egy oldNote nevu jegyzetet a repoban

	update := &vpostit.Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: vpostit.Info{
			UpdateAt: time.Now(),
		},
		ID: note.ID,
	}

	update2 := &vpostit.Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: vpostit.Info{
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

	Note1 := &vpostit.Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note2 := &vpostit.Note{
		Title: "My Today's TODO list",
		Body:  "go shopping, do housework, learn GO",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note3 := &vpostit.Note{
		Title: "My Today's Housework list",
		Body:  "make laundry, cook lunch, pick up toys, hoover everywhere",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	repo := memory.NoteRepository{}
	ctx := context.Background()
	assert.NoError(t, repo.Create(ctx, Note1))
	assert.NoError(t, repo.Create(ctx, Note2))
	assert.NoError(t, repo.Create(ctx, Note3))
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

	Note1 := &vpostit.Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note2 := &vpostit.Note{
		Title: "My Today's TODO list",
		Body:  "go shopping, do housework, learn GO",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note3 := &vpostit.Note{
		Title: "My Today's Housework list",
		Body:  "make laundry, cook lunch, pick up toys, hoover everywhere",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note4 := &vpostit.Note{
		Title: "Learning TODO",
		Body:  "maps, interfaces, RDD, TDD",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
		ID: strconv.Itoa(rand.Int()),
	}

	repo := memory.NoteRepository{}
	ctx := context.Background()
	repo.Create(ctx, Note1)
	repo.Create(ctx, Note2)
	repo.Create(ctx, Note3)

	repo.DeleteByID(ctx, Note1.ID)
	repo.DeleteByID(ctx, Note2.ID)
	repo.DeleteByID(ctx, Note3.ID)
	repo.DeleteByID(ctx, Note4.ID)
}

func TestNoteRepository_FindAllNow(t *testing.T) {

	Note1 := &vpostit.Note{
		Title: "My Today's Shopping list",
		Body:  "milk, coffee, bagels and more",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note2 := &vpostit.Note{
		Title: "My Today's TODO list",
		Body:  "go shopping, do housework, learn GO",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note3 := &vpostit.Note{
		Title: "My Today's Housework list",
		Body:  "make laundry, cook lunch, pick up toys, hoover everywhere",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	Note4 := &vpostit.Note{
		Title: "Learning TODO",
		Body:  "maps, interfaces, RDD, TDD",
		Info: vpostit.Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
		ID: strconv.Itoa(rand.Int()),
	}

	repo := memory.NoteRepository{}
	ctx := context.Background()
	repo.Create(ctx, Note1)
	repo.Create(ctx, Note2)
	repo.Create(ctx, Note3)
	repo.Create(ctx, Note4)

	notes, err := repo.FindAllNow(ctx)
	assert.NoError(t, err)
	assert.ContainExactly(t, []vpostit.Note{*Note1, *Note2, *Note3, *Note4}, notes)
}
