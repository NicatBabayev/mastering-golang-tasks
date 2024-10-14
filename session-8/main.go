package main

import (
	"fmt"
	"session-8/model/assertion"
	"session-8/model/defining"
	"session-8/model/embedding"
	"session-8/model/implementing"
	"session-8/model/switching"
)

func main() {
	// Task 1
	var c defining.Shape
	var r defining.Shape

	c = defining.Circle{Radius: 5}
	r = defining.Rectangle{
		Width:  5,
		Height: 10,
	}

	fmt.Printf("Circle area: %.14f\n", defining.AreaOfShape(c))
	fmt.Printf("Rectangle area: %.0f\n", defining.AreaOfShape(r))

	fmt.Println()

	// Task 2
	var d defining.Speaker
	var p defining.Speaker

	d = defining.Dog{
		Sound: "Woof",
	}
	p = defining.Person{
		Sound: "Hello",
	}
	fmt.Printf("Dog says: %s! \n", defining.Says(d))
	fmt.Printf("Person says: %s! \n", defining.Says(p))

	fmt.Println()

	// Task 3
	var cc implementing.PaymentProcessor
	var pl implementing.PaymentProcessor
	cc = implementing.CreditCard{Amount: 100}
	pl = implementing.Paypal{Amount: 75.5}
	value1, _ := cc.(implementing.CreditCard)
	value2, _ := pl.(implementing.Paypal)
	cc.ProcessPayment(value1.Amount)
	pl.ProcessPayment(value2.Amount)

	fmt.Println()

	//Task 4
	var emailNotifier implementing.Notifier
	var smsNotifier implementing.Notifier
	emailNotifier = implementing.EmailNotifier{Message: "Your item has shipped"}
	smsNotifier = implementing.SMSNotifier{Message: "Your item has shipped"}
	implementing.NotifyMessage(emailNotifier)
	implementing.NotifyMessage(smsNotifier)

	fmt.Println()

	// Task 5
	var i assertion.EmptyInterface = 42
	var s assertion.EmptyInterface = "GoLang"
	var f assertion.EmptyInterface = 3.1415

	assertion.TypeOfInterface(i)
	assertion.TypeOfInterface(s)
	assertion.TypeOfInterface(f)

	fmt.Println()

	// Task 6
	var i1 switching.EmptyInterface = 100
	var s1 switching.EmptyInterface = "Hi, Gophers"
	var b1 switching.EmptyInterface = true
	var f1 switching.EmptyInterface = 3.1415

	switching.TypeOfInterface(i1)
	switching.TypeOfInterface(s1)
	switching.TypeOfInterface(b1)
	switching.TypeOfInterface(f1)

	fmt.Println()

	// Task 7
	var m1 embedding.Robot
	m1 = embedding.AutoBot{}
	m1.Move()
	m1.Say()

}
