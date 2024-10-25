package main

import (
	"fmt"
	"session-10/task/buffered_unbuffered_channels"
	"session-10/task/channels_basic_operations"
	"session-10/task/goroutines"
)

func main() {
	// Task 1
	fmt.Println("// Task 1")
	goroutines.RunTask1()

	fmt.Println()

	// Task 2
	fmt.Println("// Task 2")
	goroutines.RunTask2()

	fmt.Println()

	// Task 3
	fmt.Println("// Task 3")
	channels_basic_operations.RunTask3()

	fmt.Println()

	// Task 4
	fmt.Println("// Task 4")
	channels_basic_operations.RunTask4()

	fmt.Println()

	// Task 5
	fmt.Println("// Task 5")
	buffered_unbuffered_channels.RunTask5()

	fmt.Println()

	// Task 6
	fmt.Println("// Task 6")
	buffered_unbuffered_channels.RunTask6()

}
