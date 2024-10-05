package main

import "fmt"

func main() {
	var x, y int = 10, 20
	fmt.Printf("Before swapping: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping: x = %d, y = %d\n", x, y)

}

func swap(a, b *int) {
	*a, *b = *b, *a
}
