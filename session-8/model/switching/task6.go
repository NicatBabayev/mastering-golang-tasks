package switching

import "fmt"

type EmptyInterface interface{}

func TypeOfInterface(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Type is int: %d\n", v)
	case string:
		fmt.Printf("Type is string: %s\n", v)
	case bool:
		fmt.Printf("Type is bool: %t\n", v)
	default:
		fmt.Println("Unknown type")

	}

}
