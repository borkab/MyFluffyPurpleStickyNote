package vpostit

import (
	"encoding/json"
	"fmt"
)

type ExampleStruct struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
	Baz bool   `json:"baz"`
}

func errCheck(err error) error {
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	return err
}

func JJson(ent, ent2 ExampleStruct) ExampleStruct {

	marshaledEnt, err := json.Marshal(ent) //a marshalledEnt valtozonak megadom ertekkent azt a []byte-t, ami az ent nevu struct json.Marshal func utan jott letre
	errCheck(err)
	stringedME := string(marshaledEnt) //a marhalozas utan letrejott []byte-t konvertalom stringge, es ezt a stringet megadom a stringedME valtozo ertekenek

	json.Unmarshal([]byte(stringedME), &ent2) //az unmarshal func-nek megadom parameternek az igy mar stringge alakitott []byte-t, es egy a korabban letrehozott ent2 structra mutato pointert
	errCheck(err)

	return ent2
}
