package sync_mutex

import (
	"fmt"
	"sync"
)

func incrementByOne(res *int, mux *sync.Mutex, wg *sync.WaitGroup) {
	mux.Lock()
	*res += 1
	mux.Unlock()
	defer wg.Done()
}

func RunTask5() {
	var (
		mux sync.Mutex
		wg  sync.WaitGroup
		res int
	)
	wg.Add(100)
	for range 100 {
		go incrementByOne(&res, &mux, &wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", res)

}
