package select_statement

import (
	"fmt"
	"time"
)

func RunTask1() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case data := <-ch1:
		fmt.Println("Received:", data)
	case data := <-ch2:
		fmt.Println("Received:", data)
	}

}
