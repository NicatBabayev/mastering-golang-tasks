package main

import "fmt"

func main() {

	var x int = 5
	fmt.Println("Value of x before running incrementByValue method is: ", x)
	fmt.Println("Result of method incrementByValue(x):", incrementByValue(x))
	fmt.Println("Value of x after running incrementByValue method is: ", x)

	fmt.Println()

	fmt.Println("Value of x before running incrementByPointer method is: ", x)
	res := incrementByPointer(&x)
	fmt.Println("Result of method incrementByPointer(&x):", *res)
	fmt.Println("Value of x after running incrementByValue method is: ", x)

}

func incrementByValue(val int) int {
	val += 1
	return val
}

func incrementByPointer(ptr *int) *int {
	*ptr += 1
	return ptr

}
