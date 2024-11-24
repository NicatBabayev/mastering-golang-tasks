package using_reflection

import (
	"fmt"
	"reflect"
)

type Person1 struct {
	Name string
	Age  int
}

func (p *Person1) Greet() string {
	return fmt.Sprintf("Hello, I am %s", p.Name)
}

func InvokeMethod(value interface{}, methodName string) {
	val := reflect.ValueOf(value)
	method := val.MethodByName(methodName)
	if method.IsValid() {
		res := method.Call([]reflect.Value{})[0].String()
		fmt.Printf("Invoked method: %s, Result: %s.\n", methodName, res)
	}
}

func RunTask4() {
	p := &Person1{Name: "Koroglu", Age: 7777}
	InvokeMethod(p, "Greet")

}
