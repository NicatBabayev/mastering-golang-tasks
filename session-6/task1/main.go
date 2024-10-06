package main

import "fmt"

func main() {
	var arr = [5]int{10, 20, 30, 40, 50}
	var sum int = 0

	// Print original array
	fmt.Println(" -   Original array: ", arr)

	// Print first element of array
	fmt.Println(" -   First element:", arr[0])

	// Print last element of array
	fmt.Println(" -   Last element:", arr[len(arr)-1])

	// Calculate sum
	for _, v := range arr {
		sum += v
	}

	// Print sum of array elements
	fmt.Println(" -   Sum:", sum)

	// Reverse array
	if len(arr) >= 2 {

		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
		}

		// Print the reversed arr
		fmt.Println(" -   Reversed array:", arr)
	} else {
		fmt.Printf(" -   Array consists only 1 element: [%d]. Reverse is not possible", arr[len(arr)-1])
	}

}
