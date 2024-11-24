package go_generics

import "fmt"

func MinMax[T int | float64](slice []T) (T, T) {
	var min, max T
	min = slice[0]
	max = slice[0]

	for _, num := range slice[1:] {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}
func RunTask6() {
	min, max := MinMax([]int{10, 20, 5, 8, 15})
	fmt.Printf("Min: %v, Max: %v\n", min, max)
}
