package sync_waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func runGoRoutine(i int, wg *sync.WaitGroup) {
	fmt.Printf("Goroutine %d starting\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Goroutine %d finished\n", i)
	wg.Done()
}

func RunTask3() {
	var wg sync.WaitGroup
	wg.Add(3)
	go runGoRoutine(1, &wg)
	go runGoRoutine(2, &wg)
	go runGoRoutine(3, &wg)

	wg.Wait()
	fmt.Println("All goroutines have completed")
}
