package reflection

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func InspectStruct(value interface{}) {
	typ := reflect.TypeOf(value)
	val := reflect.ValueOf(value)
	fieldCount := typ.NumField()
	for i := range fieldCount {
		fieldName := typ.Field(i).Name
		fieldType := typ.Field(i).Type
		fieldVal := val.Field(i)

		fmt.Printf("Field Name: %v, Type: %v, Value: %v\n", fieldName, fieldType, fieldVal)

	}
}

func RunTask2() {
	p := Person{
		Name: "Koroglu",
		Age:  7777,
	}
	InspectStruct(p)
}
