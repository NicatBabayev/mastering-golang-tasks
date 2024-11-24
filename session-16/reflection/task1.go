package reflection

import (
	"fmt"
	"reflect"
)

func RunTask1() {
	IdentifyTypeAndKind(42)
	IdentifyTypeAndKind("Hello")
	IdentifyTypeAndKind([]int{1, 2, 3})
}

func IdentifyTypeAndKind(value interface{}) {
	typ, val := reflect.TypeOf(value), reflect.ValueOf(value)
	fmt.Printf("Value: %v, ", val)
	fmt.Printf("Type: %v, ", typ)
	fmt.Printf("Kind: %v\n", val.Kind())
}
