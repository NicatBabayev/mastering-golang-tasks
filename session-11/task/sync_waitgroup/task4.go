package sync_waitgroup

import (
	"fmt"
	"sync"
)

func calculateSum(intSlice []int, wg *sync.WaitGroup) int {
	defer wg.Done()
	sum := 0
	for _, v := range intSlice {
		sum += v
	}
	return sum
}

func RunTask4() {
	var (
		wg    sync.WaitGroup
		psum1 int
		psum2 int
	)

	wg.Add(2)
	mainSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	firstPart := mainSlice[:len(mainSlice)/2]
	secondPart := mainSlice[len(mainSlice)/2:]
	go func() {
		psum1 = calculateSum(firstPart, &wg)

	}()
	go func() {
		psum2 = calculateSum(secondPart, &wg)
	}()

	wg.Wait()

	fmt.Printf("Partial sum 1: %d\n", psum1)
	fmt.Printf("Partial sum 2: %d\n", psum2)
	fmt.Printf("Total sum: %d\n", psum1+psum2)

}
