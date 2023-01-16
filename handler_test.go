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

	//inditok egy uj kerest, amiben megadom, h milyen metodust hivok meg rajtat, milyen cimen
	request, err := http.NewRequest(http.MethodPost, server.URL, nil)
	assert.NoError(t, err)

	response, err := server.Client().Do(request) //a kliensunk ami beolvassa a kerest
	assert.NoError(t, err)
	assert.Equal(t, http.StatusTeapot, response.StatusCode) //osszehasonlitjuk a kert Teapot statuszkodot a kapott statuszkoddal
	assert.Equal(t, "bar", response.Header.Get("X-Foo"))    //megnezem h "bar" - e a header mapben az "X-Foo" key ertekparja

	bs, err := io.ReadAll(response.Body) //beolvasom a response Body tartalmat
	assert.NoError(t, err)
	assert.NotEmpty(t, bs) //ellenorzom, h a response Body nem ures-e
	assert.NoError(t, response.Body.Close())
	assert.Contain(t, string(bs), "Hello World!\n<3") //ellenorzom, h a stringge alakitott bs valtozoba beolvasott response Body tartalmazza-e a kert szoveget
}

func TestBuzz(t *testing.T) {

	handler := &BuzzLightyearsLaserHandLER{} //a handler valtozonak megadjuk a mi handlerunket
	server := httptest.NewServer(handler)    //bedobjuk a handlerunket a serverbe
	defer server.Close()                     // ha minden kesz, bezarjuk a servert

	const expected = "To the infinity and beyond!\n:)"
	requestBody := strings.NewReader(expected) //beolvassuk egy Readerrel a string-unket amit a request Bodyba szeretnenk bedobni

	request, err := http.NewRequest(http.MethodPost, server.URL, requestBody) //inditunk egy uj kerest POST metodussal
	assert.NoError(t, err)                                                    //errCheck

	response, err := server.Client().Do(request) //a kliensunk
	fmt.Println(response)                        //kikopjuk a kijelzore a valaszt
	assert.NoError(t, err)                       //errCheck

	assert.Equal(t, http.StatusTooEarly, response.StatusCode) //megnezzuk h a vart status codeot kaptuk e
	assert.Equal(t, "foo", response.Header.Get("POST"))       //szuksegem van erre?

	request, err = http.NewRequest(http.MethodGet, server.URL, nil) //inditok egy uj kerest, GET metodussal
	assert.NoError(t, err)                                          //errCheck

	response, err = server.Client().Do(request) //kliens
	assert.NoError(t, err)                      //errCheck

	assert.Equal(t, http.StatusGone, response.StatusCode) //megnezem h a vart statuscodeot kaptam e vissza
	assert.Equal(t, "bar", response.Header.Get("GET"))    //kell ez?

	bs, err := io.ReadAll(response.Body)                     //beolvasom a response Body tartalmat
	assert.NoError(t, err)                                   //errCheck
	assert.Equal(t, strings.TrimSpace(string(bs)), expected) //megegyezik a response Body tartalma az elvart szoveggel?
}
