package main

import (
	"net/http"
	vpostit "vPOST-it"
)

func main() {

	http.ListenAndServe(":8080", &vpostit.FluffyHandler{})
	//The handler argument is valid because it implements ServeHTTP method which means
	//it implements the Handler interface which is the type of argument.

}
