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
	res.Header().Add("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusTeapot)
	res.Write([]byte("Hello World\n"))

	//res.WriteHeader(http.StatusTeapot) //dobja vissza ezt a valasz kodot: I am a tea pot

	/*
		if req.Method == "GET" { //ha a GET metodust hivom meg rajta
			res.WriteHeader(http.StatusOK) //dobja vissza ezt a statusz kodot
			res.Write(f.Fluff)             //dobja vissza valaszkent a handlerunkre mutato pointer struct Fluff fieldjenek tartalmat
		}
	*/
}
