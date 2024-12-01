package unit_testing

import (
	"reflect"
	"testing"
)

func TestNewUserProfile_Name(t *testing.T) {
	if res := NewUserProfile("Koroglu", 7777); reflect.TypeOf(res.name).Kind() != reflect.String {
		t.Error("Name field should be string")
	}
}
func TestNewUserProfile_Age(t *testing.T) {
	if res := NewUserProfile("Koroglu", 7777); reflect.TypeOf(res.age).Kind() != reflect.Int {
		t.Error("Age field should be integer")
	}
}
func TestNewUserProfile_IsMinor(t *testing.T) {
	if res := NewUserProfile("Koroglu", 17); !res.isMinor {
		t.Errorf("At age %d %s should be minor", res.age, res.name)
	}
}
