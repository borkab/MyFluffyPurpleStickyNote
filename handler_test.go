package vpostit

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adamluzsi/testcase/assert"
)

func TestHandler(t *testing.T) {
	handler := &FluffyHandler{}           //az en handler structom
	server := httptest.NewServer(handler) //nyitok egy servert es beledobom a handleremet
	defer server.Close()                  //ha minden kesz, bezarom a servert

	//inditok egy uj kerest, amiben megadom, h milyen metodust- hivok meg rajta, milyen cimen es mi legyen a tartalma
	req, err := http.NewRequest(http.MethodPost, server.URL, nil)
	assert.NoError(t, err)

	response, err := server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusTeapot, response.StatusCode)
	assert.Equal(t, "bar", response.Header.Get("X-Foo"))

	bs, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, bs)
	assert.NoError(t, response.Body.Close())
	assert.Contain(t, string(bs), "Hello World! <3")
}
