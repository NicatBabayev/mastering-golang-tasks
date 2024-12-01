package mocking

import (
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

func NotifyUser(n Notifier, message string) error {
	if message == "success" {
		return nil
	} else if message == "fail" {
		return fmt.Errorf("failed to send message")
	} else {
		return fmt.Errorf("failed to do everything")
	}
}

func RunTask8() {
	n := new(Notifier)
	fmt.Println(NotifyUser(*n, "success"))
}
