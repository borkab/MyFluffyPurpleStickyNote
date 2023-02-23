package vpostit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adamluzsi/testcase/assert"
)

//ha meghivom a Get=et a / - en akkor
//ha a repositoryban van valami akkor azt add vissza mindet
//ha nincs akkor meg semmi
//listazd a viselkedeseket, apro lepeseket, amiket meg tudsz irni
//ezekre megirod a teszteket
//aztan irsz ra kodot

/*
GET /notes
return all notes
return status code 200
return a json encoded list of note DTO value that represents a list of Note entity
*/

type repo Repository

func TestPostitHandler(t *testing.T) {
	handler := PostitHandler{}            //az en handler structom
	server := httptest.NewServer(handler) //nyitok egy servert es beledobom a handleremet
	defer server.Close()

	//when I request the GET method with / path, and no notes located in the repository, then I should get back an empty list from the server
	t.Run("call GET with / on empty repo and get an empty list", func(t *testing.T) {

		//inditok egy uj kerest ahol megadom neki h milyen methodust hivok meg rajta
		request, err := http.NewRequest(http.MethodPost, server.URL+"/", nil)
		assert.NoError(t, err)

		//a fenti keresre a kliens ad egy valszt
		response, err := server.Client().Do(request)
		assert.NoError(t, err)

		//add vissza az ures listat ha a repository ures
		want:= 

	})
	//when I request the GET method with / path, and notes are located in the repository, then I should get back all the post it note DTOs from the server
	t.Run("call GET with / on not empty repo and get all the DTOs", func(t *testing.T) {

	})
	//when I request the GET method with / path, I should get back all the post it notes from the server
	t.Run("call GET with / on not empty repo and get all the notes", func(t *testing.T) {

	})
	//when I request the GET method with /:id path, where the :id is an id that points to a non existing note, I get back 404 not found
	t.Run("call GET with /:id (unknown id), get 404", func(t *testing.T) {

	})
	//when I request the GET method with /:id path, where the :id is an id that points to an existing note, I get back the note DTO
	t.Run("call GET with /:id (known id), get the DTO", func(t *testing.T) {

	})

	t.Run("check the statuscode", func(t *testing.T) {

		request, err := http.NewRequest(http.MethodPost, server.URL+"/", nil)
		assert.NoError(t, err)

		//a fenti keresre a kliens ad egy valszt
		response, err := server.Client().Do(request)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.StatusCode) //megnezem h a vart statuscodeot kaptam e vissza
	})
}
