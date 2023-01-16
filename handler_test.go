package vpostit

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adamluzsi/testcase/assert"
)

func TestHandler(t *testing.T) {

	handler := &FluffyHandler{}           //az en handler structom
	server := httptest.NewServer(handler) //nyitok egy servert es beledobom a handleremet
	defer server.Close()                  //ha minden kesz, bezarom a servert

	const expected = "Hello World! <3"         //ezt szeretnem a request Bodyban
	requestURL := server.URL + "/hello"        //a port URL cime amin a server figyel /hello-val a vegen
	requestBody := strings.NewReader(expected) //megadom h a request Body beolvassa az elvart szoveget

	//inditok egy uj kerest, amiben megadom, h milyen metodust- hivok meg rajta, milyen cimen es mi legyen a tartalma
	req, err := http.NewRequest(http.MethodPost, requestURL, requestBody)
	assert.NoError(t, err)

	//
	resp, err := http.DefaultClient.Do(req) //ez az alap kliensunk, aminek megadjuk a keresunket
	fmt.Println(resp)                       //kiiratjuk a valaszt(response obbjektum)
	assert.NoError(t, err)

	if http.StatusTeapot != resp.StatusCode {
		t.Fatal("created code doesn't as expected. it should be: 418")
	} //ez az elvarasunk a servertol

	//req, err = http.NewRequest(http.MethodGet, requestURL, nil)
	//assert.NoError(t, err)

	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)

	if http.StatusOK != resp.StatusCode {
		t.Fatal("created code doesn't as expected. it should be: 201")
	}
	bs, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	if strings.TrimSpace(string(bs)) != expected {
		t.Fatal("oh fluff it! :( ")
	}

}
