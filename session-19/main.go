package main

import (
	"fmt"
	"session-19/task/sql_intro"
	"time"
)

func main() {
	fmt.Println("Starting the container")
	sql_intro.RunContainer()

	// Task1
	fmt.Println("// Task1")
	sql_intro.RunTask1()

	fmt.Println()

	// Task2
	fmt.Println("// Task2")
	sql_intro.RunTask2()

	defer func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Removing the container")
		sql_intro.RemoveContainer()
	}()

}
