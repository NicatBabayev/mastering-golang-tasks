package main

import "fmt"

func main() {

	var x int = 10

	var ptr *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", ptr)
	fmt.Println("Value at pointer:", *ptr)

}
