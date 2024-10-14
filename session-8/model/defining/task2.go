package defining

type Speaker interface {
	Speak() string
}

type Dog struct {
	Sound string
}
type Person struct {
	Sound string
}

func (d Dog) Speak() string {
	return d.Sound
}

func (p Person) Speak() string {
	return p.Sound
}

func Says(s Speaker) interface{} {
	return s.Speak()

}
