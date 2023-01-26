package vpostit

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
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

func JJson(ent, ent2 ExampleStruct) (ExampleStruct, bool, error) {
	//ent := ExampleStruct{}
	//ent2 := ExampleStruct{}
	var equal bool
	MEnt, err := json.Marshal(ent)
	errCheck(err)

	err = json.Unmarshal(MEnt, &ent2)
	errCheck(err)
	if !reflect.DeepEqual(ent, ent2) {
		equal = false
		err = errors.New("schade marmelade")
	}
	return ent2, equal, err
}
