package sync_rwmutex

import (
	"fmt"
	"sync"
)

func readCounter(counter *int, rwMux *sync.RWMutex, wg *sync.WaitGroup) {
	rwMux.RLock()
	fmt.Println("Reader accessed counter:", *counter)
	rwMux.RUnlock()
	defer wg.Done()
}

func writeCounter(counter *int, rwMux *sync.RWMutex, wg *sync.WaitGroup) {
	rwMux.Lock()
	*counter += 1
	fmt.Println("Writer updated counter:", *counter)
	rwMux.Unlock()
	defer wg.Done()
}

func RunTask8() {
	var (
		counter int
		wg      sync.WaitGroup
		rwMux   sync.RWMutex
	)

	wg.Add(4)
	go readCounter(&counter, &rwMux, &wg)
	go writeCounter(&counter, &rwMux, &wg)
	go readCounter(&counter, &rwMux, &wg)
	go writeCounter(&counter, &rwMux, &wg)
	wg.Wait()
}
