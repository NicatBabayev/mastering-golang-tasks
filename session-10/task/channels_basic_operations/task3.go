package channels_basic_operations

import (
	"fmt"
	"time"
)

func sendToChannel(val int, ch chan int) {
	time.Sleep(500 * time.Millisecond)
	ch <- val
}

func RunTask3() {
	val := 42
	ch := make(chan int)
	go sendToChannel(val, ch)
	fmt.Println("Received value:", <-ch)
}
