package main

import (
	"fmt"
	"session-11/task/select_statement"
	"session-11/task/sync_atomic"
	"session-11/task/sync_mutex"
	"session-11/task/sync_rwmutex"
	"session-11/task/sync_waitgroup"
)

func main() {
	// Task 1
	fmt.Println("// Task 1")
	select_statement.RunTask1()

	fmt.Println()

	// Task 2
	fmt.Println("// Task 2")
	select_statement.RunTask2()

	fmt.Println()

	// Task 3
	fmt.Println("// Task 3")
	sync_waitgroup.RunTask3()

	fmt.Println()

	// Task 4
	fmt.Println("// Task 4")
	sync_waitgroup.RunTask4()

	fmt.Println()

	// Task 5
	fmt.Println("// Task 5")
	sync_mutex.RunTask5()

	fmt.Println()

	// Task 6
	fmt.Println("// Task 6")
	sync_mutex.RunTask6()

	fmt.Println()

	// Task 7
	fmt.Println("// Task 7")
	sync_rwmutex.RunTask7()

	fmt.Println()

	// Task 8
	fmt.Println("// Task 8")
	sync_rwmutex.RunTask8()

	fmt.Println()

	// Task 9
	fmt.Println("// Task 9")
	sync_atomic.RunTask9()

}
