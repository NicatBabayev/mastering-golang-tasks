package implementing

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type CreditCard struct {
	Amount float64
}

type Paypal struct {
	Amount float64
}

func (cc CreditCard) ProcessPayment(amount float64) {
	fmt.Printf("Processing credit card payment of %.2f\n", amount)
}

func (pl Paypal) ProcessPayment(amount float64) {
	fmt.Printf("Processing Paypal payment of %.2f\n", amount)
}
