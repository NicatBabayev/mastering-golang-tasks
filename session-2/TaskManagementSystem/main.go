package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		projectStatus = "In Progress"
		tasksCreated  = 25
		isCompleted   = false
	)

	const (
		low = iota
		medium
		high
		tasksAvailable int = 100
	)

	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is: 2024-09-18 00:00:00")
	fmt.Println("Project: Task Management System")

	fmt.Println()

	fmt.Printf("  Current project status: %s\n", strings.ToUpper(projectStatus))
	fmt.Printf("  Tasks completed: %d out of %d\n", tasksCreated, tasksAvailable)
	fmt.Printf("  Task priorities: %d-Low, %d-Medium, %d-High\n", low, medium, high)
	fmt.Printf("  Is the project completed? %s\n", strconv.FormatBool(isCompleted))

}
