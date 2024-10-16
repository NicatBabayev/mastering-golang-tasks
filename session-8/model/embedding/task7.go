package embedding

import "fmt"

type Mover interface {
	Move()
}
type Sayer interface {
	Say()
}

type Robot interface {
	Mover
	Sayer
}

type AutoBot struct{}

func (a AutoBot) Move() {
	fmt.Println("Moving forward")
}

func (a AutoBot) Say() {
	fmt.Println("I am a robot")
}
