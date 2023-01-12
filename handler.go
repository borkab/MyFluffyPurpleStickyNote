//csinálj egy hello world ot válaszoló http handlert _/
//a válasz body tartalma legyen hello world,
//a válasz http kód legyen http I'm a tea pot (418)

package vpostit

import (
	"fmt"
	"io"
	"net/http"
)

type FluffyHandler struct { //create a handler struct
	Fluff []byte
}

// implement `ServeHTTP` method on `FluffyHandler` struct
func (f *FluffyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	/*
		//create response binary data
		data := &FluffyHandler{
			Fluff: []byte("Hello, World!"),
		}

		//write `data` to response
		res.Write(data.Fluff)
	*/

	//to see which methods can write a string to a Writer

	//write `Hello` using `io.WriteString` function
	io.WriteString(res, "Hello ")

	//write `World` using `fmt.Fprint` function
	fmt.Fprint(res, "World! ")

	//write `<3` using simple `Write` call
	res.Write([]byte("<3"))
}
