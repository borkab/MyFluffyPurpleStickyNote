//csinálj egy hello world ot válaszoló http handlert _/
//a válasz body tartalma legyen hello world,
//a válasz http kód legyen http I'm a tea pot (418)

package vpostit

import (
	"io"
	"net/http"
)

type FluffyHandler struct { //create a handler struct
	Fluff []byte
}

// implement `ServeHTTP` method on `FluffyHandler` struct
func (f *FluffyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	//create response binary data
	data := []byte("Hello World! <3")

	//write `data` to response
	res.Write(data)

	if req.Method == "POST" { //ha a POST metodust hivom meg rajta,
		bs, err := io.ReadAll(req.Body) //akkor olvassa be a keres Body tartalmat
		if err != nil {                 //ha vmi hiba van, kuldje vissza
			res.WriteHeader(http.StatusBadRequest) //ezt a statusz kodot
			return                                 //.
		}
		f.Fluff = bs         //a handlerunkre mutato pointer strukturank Fluff fieldjenek tartalma egyezzen meg a bs bajt slicesszal
		res.WriteHeader(418) //dobja vissza ezt a valasz kodot: I am a tea pot
	}
	if req.Method == "GET" { //ha a GET metodust hivom meg rajta
		res.WriteHeader(http.StatusOK) //dobja vissza ezt a statusz kodot
		res.Write(f.Fluff)             //dobja vissza valaszkent a handlerunkre mutato pointer struct Fluff fieldjenek tartalmat
	}
}
