//csinálj egy hello world ot válaszoló http handlert _/
//a válasz body tartalma legyen hello world,
//a válasz http kód legyen http I'm a tea pot (418)

package vpostit

import (
	"net/http"
)

type FluffyHandler struct { //my handler structs
	Fluff []byte
}

func (f *FluffyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := &FluffyHandler{
		Fluff: []byte("Hello, World!"),
	} //create response bunary data

	res.Write(data.Fluff)
}
