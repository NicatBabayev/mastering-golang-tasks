package unit_testing

import "testing"

func TestCalculate_Addition(t *testing.T) {
	if Addition(5, 10) != 15 {
		t.Error("Addition result is wrong.")
	}
}

func TestCalculate_Subtraction(t *testing.T) {
	if Subtraction(6, 10) != -4 {
		t.Error("Subtraction result is wrong")
	}
}
func TestCalculate_Multiplication(t *testing.T) {
	if Multiplication(4, 2) != 8 {
		t.Error("Multiplication result is wrong")
	}
}
func TestCalculate_Division(t *testing.T) {
	if _, err := Division(5, 0); err != nil {
		t.Error("Second operand can't be 0")
	}
}

func TestCalculate(t *testing.T) {
	if res, _ := Calculate(2, 3, "+"); res != 5 {
		t.Error("Calculation logic is not truecaw")
	}
	if _, err := Calculate(3, 4, "wrong_operation"); err != nil {
		t.Error("Operation is wrong.", err)
	}
}
