package user_inputs

import "fmt"

func RunTask5() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are not an adult")
	}
}
