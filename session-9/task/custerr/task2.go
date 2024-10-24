package custerr

import "fmt"

type DivisionError2 struct {
	msg string
}

func (de DivisionError2) Error() string {
	return de.msg
}

func Divide2(a, b float64) (float64, error) {
	err := DivisionError2{msg: ""}
	if b == 0 {
		err.msg = fmt.Sprintf("Error: Cannot divide %.0f by %.0f.\n", a, b)
		return 0.0, err
	} else {
		return a / b, nil
	}
}
