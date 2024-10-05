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
	fmt.Println("Before doubling:", one, two, three, four, five, six)
	doubleVariadicElements(&one, &two, &three, &four, &five, &six)
	fmt.Println("After doubling:", one, two, three, four, five, six)

}

func doubleVariadicElements(numbers ...*int) {
	for i := range numbers {
		*numbers[i] *= 2
	}
}
