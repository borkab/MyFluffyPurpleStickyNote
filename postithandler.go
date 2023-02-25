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

import (
	"io"
	"log"
	"net/http"
)

type PostitHandler struct{} //ez a handler

type ListOfTheNotesDTO struct { //ezek a valaszobjektumok.ebben a structban tarolom az osszes noteok listajat
	TitlesOfTheNotes []string `json:"titlesOfTheNotes"`
}

// ez egy api endpoint programoknak
func (ph PostitHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	/*
		GET /notes
		return all notes
		return status code 200
		return a json encoded list of note DTO value that represents a list of Note entity
	*/

	//1. ha a get metodussal hivom, akkor adja vissza az osszes mar letezo note-ot
	//1.1 ez egy if szerk. amiben megadom neki a feltetelt.
	//1.2 ha ez a feltetel tesjesul akkor megadom neki h mit csinaljon: irja be a request Bodyba az osszes note-ot
	//1.3 error handling: ha vmi hibara fut akkor irja ki a hibat
	if rq.Method == "GET" {

		_, err := rw.Write([]byte(InMemoryNoteRepository.MyNotes[]))
		if err != nil {
			log.Println("error", err.Error())
			return
		}
		rw.WriteHeader(http.StatusOK)
	}
	// a tesztben csinalni kell olyan lehetoseget, h ures a repo es nem tud visszaadni semmit, h akkor mit csinaljon, meg ha van benne akkor azt adja vissza..
	//2. ha ez a keres sikeresen lefutott adja vissza a statusOK statuszkodot
	//2.1 ezt is bele kell irni a fenti if szerkezetbe
	//2.2 itt is kell a tesztben  error handling: megegyezik e a megkapott statuszkod azzal amit akarok(200)
	//3. adja vissza a note DTO ertekek json-ban enkodolt listajat ami egy Note entitasok listaja
	//3.1 a tesztben csinalnom kell nehany note-ot hogy legyen mibol visszaadnia
	//3.2 json.Marshal()

}
