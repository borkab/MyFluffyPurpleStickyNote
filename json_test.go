package vpostit

import (
	"github.com/adamluzsi/testcase/assert"
	"testing"
)

type ExampleStruct struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
	Baz bool   `json:"baz"`
}

func Test_json(t *testing.T) {
	entity1 := ExampleStruct{
		Foo: "Hello, world!",
		Bar: 42,
		Baz: !false, // it's funny because it's true
	}

	// use json.Marshal with entity1

	// use json.Unmarshal on the results of the json.Marshal, and populate entity2
	var entity2 ExampleStruct

	assert.Equal(t, entity1, entity2)
}
