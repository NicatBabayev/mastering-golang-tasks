package main

import (
	"fmt"
	"strings"
)

const (
	_   = 0
	low = iota
	medium
	high
	tasksAvailable int = 100
)

func main() {
	var (
		projectStatus   = "IN PROGRESS"
		tasksCompleted  = 100
		isCompleted     = false
		tasksRemaining  = tasksAvailable - tasksCompleted
		projectProgress = ""
	)

	if tasksCompleted == tasksAvailable {
		isCompleted = true
	}

	if tasksRemaining == 0 {
		projectStatus = "FINISHED"
	}

	switch {
	case tasksCompleted < 30:
		projectProgress = "Starting phase"
	case tasksCompleted >= 30 && tasksCompleted < 60:
		projectProgress = "Midway"
	case tasksCompleted > 60:
		projectProgress = "Final phase"
	}

	fmt.Println("  Welcome to the Task Management System!")
	fmt.Println("  Project start date is: 2024-09-18 00:00:00 +0000 UTC")
	fmt.Println("  Project: Task Management System")

	fmt.Println()

	fmt.Printf("  Tasks remaining %d out of %d\n", tasksRemaining, tasksAvailable)
	fmt.Printf("  Current project status: %s\n", projectStatus)
	fmt.Printf("  Project is in the %s.\n", strings.ToLower(projectProgress))
	fmt.Printf("  Task priorities: %d-Low, %d-Medium, %d-High\n", low, medium, high)
	if tasksRemaining > 0 {
		fmt.Println("  Task list: ")
		for i := 1; i <= tasksRemaining; i++ {
			fmt.Println("  - Task", i)
		}
	} else {
		fmt.Println("  There are no remaining tasks left ")
	}

	fmt.Printf("  Is the project completed? %v\n", isCompleted)

}
