package main

import "fmt"

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Print array before manipulations
	fmt.Println("Original array before manipulations:", arr)

	// Extract slice from index 2:6
	slc := arr[2:6]
	fmt.Println("Sub-slice:", slc)

	// Append [10,11,12] to slc
	slc = append(slc, 10, 11, 12)
	fmt.Println("Appended slice:", slc)

	// Show effect of append()
	fmt.Println("Original array after manipulations:", arr)
	fmt.Println("Print original array: modified based on capacity:", slc)

}
