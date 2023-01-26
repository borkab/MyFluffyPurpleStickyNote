package vpostit

import (
	"encoding/json"
	"testing"

	"github.com/adamluzsi/testcase/assert"
)

func Test_json(t *testing.T) {
	entity1 := ExampleStruct{
		Foo: "Hello, world!",
		Bar: 42,
		Baz: !false, // it's funny because it's true
	}

	// use json.Marshal with entity1
	marshalledEnt1, err := json.Marshal(entity1)
	assert.NoError(t, err)
	// use json.Unmarshal on the results of the json.Marshal, and populate entity2
	var entity2 ExampleStruct
	json.Unmarshal(marshalledEnt1, &entity2)
	assert.Equal(t, entity1, entity2)
}

func TestJJson(t *testing.T) {
	entity1 := ExampleStruct{
		Foo: "Hello, world!",
		Bar: 42,
		Baz: !false, // it's funny because it's true
	}
	var entity2 ExampleStruct
	_, _, err := JJson(entity1, entity2)
	assert.Equal(t, entity1, entity2)
	assert.NoError(t, err)
}
