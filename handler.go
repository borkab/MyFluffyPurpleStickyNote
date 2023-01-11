package vpostit

import (
	"fmt"
	"net/http"
)

//csinálj egy hello world ot válaszoló http handlert
//a válasz body tartalma legyen hello world,
//a válasz http kód legyen http I'm a tea pot (418)

func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, World!")
}
