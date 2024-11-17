package unit_testing

import "fmt"

type UserProfile struct {
	name    string
	age     int
	isMinor bool
}

func NewUserProfile(name string, age int) UserProfile {
	var isMinor bool
	if age < 18 {
		isMinor = true
	}
	return UserProfile{
		name:    name,
		age:     age,
		isMinor: isMinor,
	}
}

func RunTask2() {
	fmt.Println(NewUserProfile("Koroglu", 7777))
}
