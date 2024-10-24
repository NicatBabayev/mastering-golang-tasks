package buffered_unbuffered_channels

import (
	"fmt"
	"time"
)

func RunTask5() {
	ch := make(chan int, 3)
	fmt.Println("Sent values into buffered channel")
	ch <- 10
	ch <- 20
	ch <- 30
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("All values received")

}
