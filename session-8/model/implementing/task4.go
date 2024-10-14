package implementing

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	Message string
}

type SMSNotifier struct {
	Message string
}

func (en EmailNotifier) Notify(message string) {
	fmt.Printf("Sending email notification: %s\n", message)
}

func (sn SMSNotifier) Notify(message string) {
	fmt.Printf("Sending SMS notification: %s\n", message)
}

func NotifyMessage(n Notifier) {
	if v, ok := n.(EmailNotifier); ok {
		n.Notify(v.Message)
	} else if v, ok := n.(SMSNotifier); ok {
		n.Notify(v.Message)
	}

}
