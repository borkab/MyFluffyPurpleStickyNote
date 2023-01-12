package vpostit

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func ErrCheck(err error) {
	if err == nil {
		log.Fatal(err)
	}
}
func TestHandler(t *testing.T) {

	handler := &FluffyHandler{}           //az en handler structom
	server := httptest.NewServer(handler) //nyitok egy servert es beledobom a handleremet
	defer server.Close()                  //ha minden kesz, bezarom a servert

	t.Run("", func(t *testing.T) {
		const expected = "Hello World! <3"         //ezt szeretnem a request Bodyban
		requestURL := server.URL + "/hello"        //a port URL cime amin a server figyel /hello-val a vegen
		requestBody := strings.NewReader(expected) //megadom h a request Body beolvassa az elvart szoveget

		//inditok egy uj kerest, amiben megadom, h milyen metodust- hivok meg rajta, milyen cimen es mi legyen a tartalma
		req, err := http.NewRequest(http.MethodPost, requestURL, requestBody)
		ErrCheck(err)

		//
		resp, err := http.DefaultClient.Do(req) //ez az alap kliensunk, aminek megadjuk a keresunket
		fmt.Println(resp)                       //kiiratjuk a valaszt(response obbjektum)
		ErrCheck(err)

		if http.StatusCreated != resp.StatusCode {
			t.Fatal("created code doen't as expected. it should be: 201")
		} //ez az elvarasunk a servertol

		req, err = http.NewRequest(http.MethodGet, requestURL, nil)
		ErrCheck(err)

		resp, err = http.DefaultClient.Do(req)
		ErrCheck(err)

		if http.StatusOK != resp.StatusCode {
			t.Fatal("created code doesn't as expected. it should be: 201")
		}
		bs, err := io.ReadAll(resp.Body)
		ErrCheck(err)

		if strings.TrimSpace(string(bs)) != expected {
			t.Fatal("oh fluff it! :( ")
		}

	})
}
