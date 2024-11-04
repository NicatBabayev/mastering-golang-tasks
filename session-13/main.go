package main

import (
	"fmt"
	"session-13/task/json"
	"session-13/task/reading_writing"
	"session-13/task/xml"
)

func main() {
	// Task 1
	fmt.Println("// Task 1 ")
	reading_writing.RunTask1()

	fmt.Println()

	// Task 2
	fmt.Println("// Task 2 ")
	reading_writing.RunTask2()

	fmt.Println()

	// Task 3
	fmt.Println("// Task 3")
	json.RunTask3()

	fmt.Println()

	// Task 4
	fmt.Println("// Task 4")
	json.RunTask4()

	fmt.Println()

	// Task 5
	fmt.Println("// Task 5")
	xml.RunTask5()

	fmt.Println()

	// Task 6
	fmt.Println("// Task 6")
	xml.RunTask6()

}
