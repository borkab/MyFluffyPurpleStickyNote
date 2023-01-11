package main

import (
	"net/http"
	vpostit "vPOST-it"
)

func main() {

	http.HandleFunc("/hello", vpostit.Handler)
	http.ListenAndServe(":8080", nil)
}
