package sync_rwmutex

import (
	"fmt"
	"sync"
)

func readFromMap(name string, gradesMap map[string]int, rwMux *sync.RWMutex, wg *sync.WaitGroup) {
	rwMux.RLock()
	_ = gradesMap[name]
	// fmt.Println("Just read from map: ", name, gradesMap[name]) // Just for testing the code
	rwMux.RUnlock()
	defer wg.Done()

}
func writeToMap(name string, grade int, gradesMap map[string]int, rwMux *sync.RWMutex, wg *sync.WaitGroup) {
	rwMux.Lock()
	gradesMap[name] = grade
	rwMux.Unlock()
	defer wg.Done()
}

func RunTask7() {
	var (
		gradesMap map[string]int
		wg        sync.WaitGroup
		rwMux     sync.RWMutex
	)
	gradesMap = make(map[string]int)
	wg.Add(6)

	go writeToMap("Garay", 20, gradesMap, &rwMux, &wg)
	go writeToMap("Ali", 25, gradesMap, &rwMux, &wg)
	go writeToMap("Medina", 28, gradesMap, &rwMux, &wg)

	go readFromMap("Garay", gradesMap, &rwMux, &wg)
	go readFromMap("Ali", gradesMap, &rwMux, &wg)
	go readFromMap("Medina", gradesMap, &rwMux, &wg)

	wg.Wait()
	fmt.Println("User data", gradesMap)
}
