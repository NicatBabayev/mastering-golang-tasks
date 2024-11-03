package ctx_pkg

import (
	"context"
	"fmt"
	"time"
)

func longRunningProgram() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func cancelFuncTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), (3*time.Second)-(time.Millisecond)) // To achieve consistent results across executions.
	defer cancel()
	go longRunningProgram()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("(cancellation)")
			return
		}
	}
}

func RunTask1() {
	cancelFuncTimeout()
}
