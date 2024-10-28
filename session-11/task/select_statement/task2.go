package select_statement

import (
	"fmt"
	"time"
)

func waitOnChannel(ch chan string) {
	select {
	case data := <-ch:
		fmt.Printf("Received data: %s", data)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout occurred: no message received")
	}
}

func RunTask2() {
	ch := make(chan string)
	waitOnChannel(ch)
}
