//csinálj egy hello world ot válaszoló http handlert _/
//a válasz body tartalma legyen hello world,
//a válasz http kód legyen http I'm a tea pot (418)

package vpostit

import (
	"net/http"
)

type FluffyHandler struct { //create a handler struct
	//Fluff []byte
}

// implement `ServeHTTP` method on `FluffyHandler` struct
func (f *FluffyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("X-Foo", "bar")
	res.WriteHeader(http.StatusTeapot)
	_, _ = res.Write([]byte("Hello World\n"))
}
