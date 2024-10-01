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
		tasksCompleted  = 92
		projectProgress = ""
	)

	switch {
	case tasksCompleted < 30:
		projectProgress = "Starting phase"
	case tasksCompleted >= 30 && tasksCompleted <= 60:
		projectProgress = "Midway"
	case tasksCompleted > 60:
		projectProgress = "Final phase"
	}

	fmt.Println("  Welcome to the Task Management System!")
	fmt.Println("  Project start date is: 2024-09-18 00:00:00 +0000 UTC")
	fmt.Println("  Project: Task Management System")

	fmt.Println()

	fmt.Printf("  Tasks remaining %d out of %d.", tasksRemaining(tasksCompleted), tasksAvailable)
	if _, f := isNearlyFinished(tasksCompleted); f {
		fmt.Println(" Project nearly finished.")
	} else {
		fmt.Println()
	}
	printProjectStatus(tasksCompleted)
	fmt.Println("  Recursive countdown of remaining tasks:")
	tasksRemainingRecursive(tasksCompleted)
	fmt.Printf("  Project is in the %s.\n", strings.ToLower(projectProgress))
	fmt.Printf("  Task priorities: %d-Low, %d-Medium, %d-High\n", low, medium, high)
	listTaskNums(tasksCompleted)

	fmt.Printf("  Is the project completed? %v\n", isCompleted(tasksCompleted, tasksAvailable))

}

func isCompleted(tasksCompleted, tasksAvailable int) bool {
	if tasksCompleted == tasksAvailable {
		return true
	} else {
		return false
	}
}
func tasksRemaining(tasksCompleted int) int {
	return tasksAvailable - tasksCompleted
}

func listTaskNums(tasksCompleted int) {
	if tasksRemaining(tasksCompleted) > 0 {
		fmt.Println("  Task list: ")
		for i := 1; i <= tasksRemaining(tasksCompleted); i++ {
			fmt.Printf("  %d: Task %d\n", i, i)
		}
	} else {
		fmt.Println("  There are no remaining tasks left ")
	}
}
func printProjectStatus(tasksCompleted int) {
	var projectStatus string

	if tasksRemaining(tasksCompleted) == 0 {
		projectStatus = "FINISHED"
	} else {
		projectStatus = "IN PROGRESS"
	}
	fmt.Printf("  Current project status: %s\n", projectStatus)
}

func isNearlyFinished(tasksCompleted int) (int, bool) {
	var isNearlyFinished bool
	if tasksCompleted > 90 {
		isNearlyFinished = true
	} else {
		isNearlyFinished = false
	}

	return tasksRemaining(tasksCompleted), isNearlyFinished
}

func tasksRemainingRecursive(tasksCompleted int) {
	t := tasksCompleted

	if t == 0 {
		fmt.Println("  All tasks completed!")
		return
	}

	fmt.Printf("  Tasks remaining: %d\n", t)
	t--

	tasksRemainingRecursive(t)
}
