/*
POST /notes
create a new note and returns the created note JSON serialized note
return status code 201
return a json encoded note DTO value that represents a Note entity
GET /notes/:id
returns a note that is referenced to an ID
return status code 200
return a json encoded note DTO value that represents a Note entity
PUT /notes/:id
allow for updating a note with new values
return status code 204 No Content
DELETE /notes/:id
delete a note in the system associated with an ID
return status code 204 No Content
*/
package vpostit

import "net/http"

type postitHandler struct{} //ez a handler

type ListOfTheNotesDTO struct { //ebben a structban tarolom az osszes noteok listajat
	TitlesOfTheNotes []string `json:"titlesOfTheNotes"`
}

func (ph postitHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	/*
		GET /notes
		return all notes
		return status code 200
		return a json encoded list of note DTO value that represents a list of Note entity
	*/

	if rq.Method == "GET"+"/notes" {
		rw.WriteHeader(http.StatusOK)
	}
}
