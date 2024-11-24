package main

import (
	"fmt"
	"session-16/go_generics"
	"session-16/reflection"
	"session-16/using_reflection"
)

func main() {
	// Task1
	fmt.Println("// Task1")
	reflection.RunTask1()

	fmt.Println()

	// Task2
	fmt.Println("// Task2")
	reflection.RunTask2()

	fmt.Println()

	// Task3
	fmt.Println("// Task3")
	using_reflection.RunTask3()

	fmt.Println()

	// Task 4
	fmt.Println("// Task4")
	using_reflection.RunTask4()

	fmt.Println()

	// Task5
	fmt.Println("// Task5")
	go_generics.RunTask5()

	fmt.Println()

	// Task6
	fmt.Println("// Task6")
	go_generics.RunTask6()
}
