package ctx_pkg

import (
	"context"
	"fmt"
	"time"
)

func sleepFiveSeconds() {
	time.Sleep(time.Second * 5)
}

func cancelWithContext() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go sleepFiveSeconds()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled due to timeout")
			return
		}
	}
}

func RunTask2() {

	cancelWithContext()

}
