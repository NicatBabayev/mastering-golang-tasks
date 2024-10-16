package assertion

import "fmt"

type EmptyInterface interface{}

func TypeOfInterface(i interface{}) {

	if v, ok := i.(int); ok {
		fmt.Printf("Value is of type int: %d\n", v)
	} else if v, ok := i.(string); ok {
		fmt.Printf("Value is of type int: %s\n", v)
	} else if v, ok := i.(float64); ok {
		fmt.Printf("Value is of type int: %f\n", v)
	}

}
