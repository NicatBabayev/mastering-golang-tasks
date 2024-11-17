package benchmarking

import (
	"fmt"
	"math/rand"
)

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	fmt.Println(arr)
}

func createRandomArray(size int) []int {
	res := make([]int, 0)
	for range size {
		randI := rand.Intn(size + 1)
		res = append(res, randI)
	}
	return res
}

func RunTask5() {
	arr := createRandomArray(10)
	bubbleSort(arr)
}
