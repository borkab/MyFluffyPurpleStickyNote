//csinálj egy hello world ot válaszoló http handlert _/
//a válasz body tartalma legyen hello world,
//a válasz http kód legyen http I'm a tea pot (418)

package vpostit

import (
	"fmt"
	"io"
	"net/http"
)

type Fluff struct {
	peggy []byte
}

func (fu *Fluff) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bs, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fu.peggy = bs
		w.WriteHeader(418)
	}
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write(fu.peggy)
	}
	fmt.Fprintf(w, "Hello, World!")
}
