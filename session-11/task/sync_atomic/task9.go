package sync_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func incrementByTen(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(counter, 10)

}

func RunTask9() {
	var (
		wg      sync.WaitGroup
		counter int64
	)
	wg.Add(10)
	for range 10 {
		go incrementByTen(&counter, &wg)
	}

	wg.Wait()
	fmt.Println("Final atomic counter value:", counter)

}
