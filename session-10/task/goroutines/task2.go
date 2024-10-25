package goroutines

import (
	"fmt"
	"time"
)

func printAlphabet() {
	for i := 65; i <= 69; i++ {
		fmt.Println(string(i))
		time.Sleep(200 * time.Millisecond)
	}
}

func printNumbers1() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}

func RunTask2() {
	go printAlphabet()
	go printNumbers1()
	fmt.Println("Main finished")
	time.Sleep(2 * time.Second)
}
