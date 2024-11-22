package main

import (
	"fmt"
	"session-14/task/benchmarking"
	"session-14/task/mocking"
	"session-14/task/table_driven"
	"session-14/task/unit_testing"
)

func main() {
	// Task1
	fmt.Println("// Task1")
	unit_testing.RunTask1()

	fmt.Println()

	// Task2
	fmt.Println("// Task2")
	unit_testing.RunTask2()

	fmt.Println()

	// Task3
	fmt.Println("// Task3")
	table_driven.RunTask3()

	fmt.Println()

	// Task4
	fmt.Println("// Task4")
	table_driven.RunTask4()

	fmt.Println()

	// Task5
	fmt.Println("// Task 5")
	benchmarking.RunTask5()

	fmt.Println()

	// Task6
	fmt.Println("// Task6")
	benchmarking.RunTask6()

	fmt.Println()

	// Task7
	fmt.Println("// Task7")
	mocking.RunTask7()

	fmt.Println()

	// Task8
	fmt.Println("// Task8")
	mocking.RunTask8()
}
