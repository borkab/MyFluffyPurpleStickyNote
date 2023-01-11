package main

//This file is in /home/borkab/learngo/github/vPOST-it/cmd/webserver, which is a nested module in the /home/borkab/learngo/github/vPOST-it/go.mod module.
//To work on multiple modules at once, please use a go.work file.

import (
	"net/http"
	vpostit "vPOST-it" //package vPOST-it is not in GOROOT (/usr/lib/go/src/vPOST-it) (compile)
)

func main() {

	http.ListenAndServe(":8080", &vpostit.Fluff{})
}
