package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		projectStatus string = "In Progress"
		tasksCreated  int    = 25
		isCompleted   bool   = false
	)

	const (
		low = iota
		medium
		high
		tasksAvailable int = 100
	)
	_ = projectStatus
	_ = tasksCreated
	_ = isCompleted
	_ = low

	fmt.Println(`  Welcome to the Task Management System!
  Project start date is: 2024-09-18 00:00:00
  Project: Task Management System`)

	fmt.Println()

	fmt.Printf("  Current project status: %s\n", projectStatus)
	fmt.Printf("  Tasks completed: %d out of %d\n", tasksCreated, tasksAvailable)
	fmt.Printf("  Task priorities: %d-Low, %d-Medium, %d-High\n", low, medium, high)
	fmt.Printf("  Is the project completed? %s\n", strconv.FormatBool(isCompleted))

}
