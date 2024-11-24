package go_generics

import "fmt"

func Sum[T int | float64](slice []T) T {
	var sum T
	for _, num := range slice {
		sum += num
	}

	return sum
}

func RunTask5() {
	fmt.Println("Sum:", Sum([]int{1, 2, 3, 4}))
	fmt.Println("Sum:", Sum([]float64{1.5, 2.5, 3.5}))
}
