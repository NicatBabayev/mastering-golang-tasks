package goroutines

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := range 5 {
		fmt.Println(i + 1)
		time.Sleep(time.Millisecond)
	}

}

func RunTask1() {
	fmt.Println("Main function started")
	go printNumbers()
	fmt.Println("Main function ended")
	time.Sleep(time.Second)
}
