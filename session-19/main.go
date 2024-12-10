package main

import (
	"fmt"
	"session-19/task/advanced"
	"session-19/task/db_sql_pkg"
)

func main() {
	// Task 1
	fmt.Println("// Task 1")
	fmt.Println("For task 1 result check the task1.png file in sql_intro directory")

	fmt.Println()

	// Task 2
	fmt.Println("// Task 2")
	fmt.Println("For task 2 result check the task2.png file in sql_intro directory")

	fmt.Println()

	// Task 3
	fmt.Println("// Task 3")
	db_sql_pkg.RunTask3()

	fmt.Println()

	// Task4
	fmt.Println("// Task 4")
	db_sql_pkg.RunTask4()

	fmt.Println()

	// Task 5
	fmt.Println("// Task 5")
	advanced.RunTask5()

	fmt.Println()

	// Task 6
	fmt.Println("// Task 6")
	advanced.RunTask6()

}
