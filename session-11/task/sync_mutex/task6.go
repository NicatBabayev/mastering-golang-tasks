package sync_mutex

import (
	"fmt"
	"sync"
)

func updateGrades(name string, grade int, gradesMap map[string]int, mux *sync.Mutex, wg *sync.WaitGroup) {
	mux.Lock()
	gradesMap[name] = grade
	mux.Unlock()
	defer wg.Done()
}

func RunTask6() {
	var (
		gradesMap map[string]int
		mux       sync.Mutex
		wg        sync.WaitGroup
	)
	gradesMap = make(map[string]int)
	wg.Add(3)
	go updateGrades("Garay", 90, gradesMap, &mux, &wg)
	go updateGrades("Ali", 85, gradesMap, &mux, &wg)
	go updateGrades("Medina", 78, gradesMap, &mux, &wg)

	wg.Wait()
	fmt.Printf("Final Grades: %v\n", gradesMap)
}
