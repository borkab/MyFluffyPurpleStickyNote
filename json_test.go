package vpostit

import (
	"encoding/json"
	"fmt"
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
	marshalledEnt1, err := json.Marshal(entity1) //marshalledEnt1 az entity1 struct marshallozasa utan letrejott []byte
	assert.NoError(t, err)
	fmt.Println(entity1)        //kiirjuk az entity1 valtozo tartalmat(ami egy ExampleStruct )
	fmt.Println(marshalledEnt1) //kiirjuk az entity1 marshalozasa soran letrejott []byte tartalmat

	// use json.Unmarshal on the results of the json.Marshal, and populate entity2
	var entity2 ExampleStruct                   //letrehozunk egy ures ExampleStructot
	fmt.Println(entity2)                        //ezt az ures ExampleStructot kiirjuk a kepernyore
	strEntity := string(marshalledEnt1)         //a strEntity valtozonak megadjuk erteknek az entity1 Examplestruct marshalozasa soran letrejott []byte tartalmat stringge konvertalva
	fmt.Println(strEntity)                      //ezt a stringge konvertalt []byte-t kiirjuk a kepernyore
	json.Unmarshal([]byte(strEntity), &entity2) //a mar stringge konvertalt byte sliceunkkat visszaalakitva stringge megadjuk az Unmarshal func argumensenek, egy az ures ExampleStruct tipusu structuncra mutato pointerrel egyutt
	fmt.Println(entity2)                        //kiirjuk a kepernyore az entity2 tartalmat ES EROSEN REMELJUK HOGY NEM URES!!!!!
	//fmt.Println(strEntity) //kiirjuk a kepernyore a stringge alakitott
	assert.Equal(t, entity1, entity2)
}

func TestJJson(t *testing.T) {
	entity1 := ExampleStruct{
		Foo: "Hello, world!",
		Bar: 42,
		Baz: !false, // it's funny because it's true
	}
	var entity2 ExampleStruct
	entity2, _, err := JJson(entity1, entity2)
	//fmt.Println(entity1)
	//fmt.Println(entity2)
	//fmt.Println(equal)
	//fmt.Println(err)
	assert.Equal(t, entity1, entity2)
	assert.NoError(t, err)
}
