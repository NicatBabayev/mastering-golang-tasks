package using_reflection

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	City string
}

func SetFieldValue(value interface{}, fieldName string, newValue interface{}) {
	val := reflect.ValueOf(value).Elem()
	fieldKind := val.FieldByName(fieldName).Type().Kind()

	switch fieldKind {
	case reflect.Int:
		val.FieldByName(fieldName).SetInt(newValue.(int64))
	case reflect.String:
		val.FieldByName(fieldName).SetString(newValue.(string))
	case reflect.Float64:
		val.FieldByName(fieldName).SetFloat(newValue.(float64))
	case reflect.Bool:
		val.FieldByName(fieldName).SetBool(newValue.(bool))
	}

}
func RunTask3() {
	p := &Person{Name: "Koroglu", Age: 7777, City: "Baku"}
	SetFieldValue(p, "City", "Cenlibel")
	fmt.Printf("%+v", p)
}
