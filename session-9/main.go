package main

import (
	"errors"
	"fmt"
	"session-9/task/custerr"
	"session-9/task/wrapper"
)

func main() {

	// Task 1
	res, err := custerr.Divide(5, 0)
	fmt.Println("// Task 1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %.1f\n", res)
	}

	fmt.Println()

	// Task 2
	res2, err2 := custerr.Divide2(5, 0)
	fmt.Println("// Task 2")
	if err2 != nil {
		fmt.Print(err2)
	} else {
		fmt.Printf("Result: %.1f", res2)
	}

	fmt.Println()

	// Task 3
	fmt.Println("// Task 3")
	err3 := wrapper.OpenFile("nonexistent.txt")
	if err3 != nil {
		fmt.Print(fmt.Errorf("Wrapped Error: %w", err3))
		fmt.Println()
		fmt.Print(fmt.Errorf("Original Error: %w", errors.Unwrap(err3)))
	}

	fmt.Println()

	// Task 4
	fmt.Println("// Task 4")
	err4 := wrapper.OpenFile2("nonexistent.txt")
	if err4 != nil {
		fmt.Print(fmt.Errorf("error: %w", err4))
	}

}
