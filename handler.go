//csinálj egy hello world ot válaszoló http handlert _/
//a válasz body tartalma legyen hello world,
//a válasz http kód legyen http I'm a tea pot (418)

package vpostit

import (
	"io"
	"net/http"
)

type FluffyHandler struct { //create a handler struct
	//Fluff []byte
}

// implement `ServeHTTP` method on `FluffyHandler` struct
func (f *FluffyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//beallitom a header map tartalmat
	res.Header().Set("X-Foo", "bar")
	//megadom h mi legyen a statuscode
	res.WriteHeader(http.StatusTeapot)
	//megadom a response Body tartalmat
	_, _ = res.Write([]byte("Hello World!\n<3"))

}

type BuzzLightyearsLaserHandLER struct {
	Buzz []byte
}

func (b *BuzzLightyearsLaserHandLER) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("POST", "foo")

	resp.Header().Add("GET", "bar")

	//_, _ = resp.Write([]byte("To the infinity and beyond!\n:)\n"))

	if req.Method == "POST" { //megjegyezzuk
		bs, err := io.ReadAll(req.Body) //beolvassuk a request Body tartalmat
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest) //ha vmi hibara fut, visszakuldjuk a hibas StatusCodeot
			return                                  //ebbol a folyamatbol visszaterunk
		}
		b.Buzz = bs           //a handler struct-unk Buzz fieldjenek megadjuk a beolvasott request Body tartalmat
		resp.WriteHeader(425) // visszadobjuk a TooEarly StatusCodot
	}
	if req.Method == "GET" { //ha a "GET" metodust hivtuk meg a handlerunkon
		resp.WriteHeader(410) //visszadobjuk a Gone Statuscodeot
		resp.Write(b.Buzz)    //megadjuk request Bodynak a Handlerunk Buzz fieldjenek a tartalmat
	}
}
