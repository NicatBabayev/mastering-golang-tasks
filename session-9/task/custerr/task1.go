package custerr

type DivisionError struct {
	msg string
}

func (de DivisionError) Error() string {
	return de.msg
}

func Divide(a, b float64) (float64, error) {
	err := DivisionError{msg: "Error: Division by zero is not allowed."}
	if b == 0 {
		return 0.0, err
	} else {
		return a / b, nil
	}
}
