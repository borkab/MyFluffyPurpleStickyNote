package vpostit

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adamluzsi/testcase/assert"
)

func TestHandler(t *testing.T) {

	handler := FluffyHandler{}            //az en handler structom
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

	// csinálj egy tesztet ahol egy request headerek közül mondjuk az X-Foo header
	// értéket valaszolod viszza a response body ban
	expected := response.Header.Get("X-Foo") //egy “expected” valtozoba bele teszem a http keresemre kapott valasz headerjeibol az X-Foo headert.
	assert.Equal(t, expected, string(bs))    //Ezt aztan ossze hasonlitom a http keresemre kapott valasz body -jabol kiolvasott tartalommal.

	var resp http.ResponseWriter
	header := request.Header.Get("X=Foo")
	_, err = resp.Write([]byte(header))
	assert.NoError(t, err)
	requestBody := strings.NewReader(header)
	req, err := http.NewRequest(http.MethodPost, server.URL, requestBody)

}

// handler ami képes mást válaszolni annak függvényében hogy GET vagy POST amit hívtál rajta
// ha gettel hívod akkor a válasz foo + code 200
// ha posttal akkor meg bar + code 201
func TestBuzz_smoke(t *testing.T) {
	handler := BuzzLightyearsLaserHandLER{}

	t.Run("on GET", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "foo\n", response.Body.String())
	})

	t.Run("on POST", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, "bar\n", response.Body.String())
	})
}

func TestBuzz(t *testing.T) {

	handler := &BuzzLightyearsLaserHandLER{} //a handler valtozonak megadjuk a mi handlerunket
	server := httptest.NewServer(handler)    //bedobjuk a handlerunket a serverbe
	defer server.Close()                     // ha minden kesz, bezarjuk a servert

	request, err := http.NewRequest(http.MethodPost, server.URL, nil) //inditunk egy uj kerest POST metodussal
	assert.NoError(t, err)                                            //errCheck

	response, err := server.Client().Do(request) //a kliensunk
	assert.NoError(t, err)                       //errCheck

	assert.Equal(t, http.StatusOK, response.StatusCode) //megnezzuk h a vart status codeot kaptuk e
	assert.Equal(t, "bar", response.Header.Get("POST")) //szuksegem van erre?

	request, err = http.NewRequest(http.MethodGet, server.URL, nil) //inditok egy uj kerest, GET metodussal
	assert.NoError(t, err)                                          //errCheck

	response, err = server.Client().Do(request) //kliens
	assert.NoError(t, err)                      //errCheck

	assert.Equal(t, http.StatusOK, response.StatusCode) //megnezem h a vart statuscodeot kaptam e vissza
	assert.Equal(t, "foo", response.Header.Get("GET"))  //kell ez?
}
