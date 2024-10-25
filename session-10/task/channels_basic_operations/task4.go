package channels_basic_operations

import "fmt"

func sendToChannel1(val []int, ch chan []int) {
	ch <- val
}

func RunTask4() {
	ch := make(chan []int)
	defer func() {
		close(ch)
		fmt.Println("Channel closed")
	}()
	go sendToChannel1([]int{1, 2, 3, 4, 5}, ch)
	numbers := <-ch
        for _, v := range numbers {
		fmt.Println(v)
	}
}
