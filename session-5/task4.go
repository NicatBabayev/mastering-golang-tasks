package main

import (
	"fmt"
)

func main() {
	var (
		one   = 1
		two   = 2
		three = 3
		four  = 4
		five  = 5
		six   = 6
	)
	doubleVariadicElements(&one, &two, &three, &four, &five, &six)

}

func convertToNum(numbers ...*int) []int {
	var arr []int
	for i := range numbers {
		arr = append(arr, *numbers[i])
	}
	return arr
}

func doubleVariadicElements(numbers ...*int) {
	fmt.Println("Before doubling:", convertToNum(numbers...))
	for i := range numbers {
		*numbers[i] *= 2
	}
	fmt.Println("After doubling:", convertToNum(numbers...))
}
