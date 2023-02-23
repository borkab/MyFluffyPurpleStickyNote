package main

import (
	"net/http"
	vpostit "vPOST-it"
)

func main() {

	http.ListenAndServe(":8080", &vpostit.PostitHandler{})
}
