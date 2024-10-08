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
		tasksCompleted  = 90
		projectProgress = ""
	)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  Recovered from panic: Error: tasksCompleted exceeds total tasks")
			fmt.Println("  Status check completed.")
		}
	}()

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
	if c, _ := calculateTasksRemaining(tasksCompleted); c >= 0 {
		fmt.Printf("  Tasks remaining %d out of %d.", tasksRemaining(tasksCompleted), tasksAvailable)
		if _, f := calculateTasksRemaining(tasksCompleted); f {
			fmt.Println(" Project nearly finished.")
		} else {
			fmt.Println()
		}
		printProjectStatus(tasksCompleted)
		if c, _ := calculateTasksRemaining(tasksCompleted); c != 0 {
			fmt.Printf("  Project is in the %s.\n", strings.ToLower(projectProgress))
		}
		fmt.Printf("  Task priorities: %d-Low, %d-Medium, %d-High\n", low, medium, high)
		listTaskNums(tasksCompleted)
		if c, _ := calculateTasksRemaining(tasksCompleted); c != 0 {
			fmt.Println("  Recursive countdown of remaining tasks:")
		}
		tasksRemainingRecursive(tasksCompleted)
		statusCheck(tasksCompleted)

		fmt.Printf("  Is the project completed? %v\n", isCompleted(tasksCompleted, tasksAvailable))

	} else {
		statusCheck(tasksCompleted)
		fmt.Printf("  Is the project completed? %v\n", isCompleted(tasksCompleted, tasksAvailable))
	}
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
	if c, _ := calculateTasksRemaining(tasksCompleted); c > 0 {
		fmt.Println("  Task list: ")
		for i := 1; i <= tasksCompleted; i++ {
			fmt.Printf("  %d: Task %d\n", i, i)
		}
	} else {
		fmt.Println("  There are no remaining tasks left ")
	}
}
func printProjectStatus(tasksCompleted int) {
	var projectStatus string

	if c, _ := calculateTasksRemaining(tasksCompleted); c == 0 {
		projectStatus = "FINISHED"
	} else {
		projectStatus = "IN PROGRESS"
	}
	fmt.Printf("  Current project status: %s\n", projectStatus)
}

func calculateTasksRemaining(tasksCompleted int) (int, bool) {
	tasksRemaining := tasksAvailable - tasksCompleted
	nearCompletion := tasksCompleted > 90
	return tasksRemaining, nearCompletion
}

func tasksRemainingRecursive(tasksCompleted int) {
	t := tasksCompleted

	if t == 0 || t == tasksAvailable {
		fmt.Println("  All tasks completed!")
		return
	}

	fmt.Printf("  Tasks remaining: %d\n", tasksAvailable-t)
	t++

	tasksRemainingRecursive(t)
}

// Variadic function to print any number of tasks
func printTasks(tasks ...string) {
	fmt.Println("Task list:")
	for i, task := range tasks {
		fmt.Printf("%d: %s\n", i+1, task)
	}
}

func statusCheck(tasksCompleted int) {
	if tasksAvailable-tasksCompleted < 0 {
		panic("  Number of completed tasks can't be more than available tasks")
	}
}
