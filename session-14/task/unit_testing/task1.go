package unit_testing

import "fmt"

func RunTask1() {
	fmt.Println(Calculate(5, 10, "-"))
}

func Addition(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Multiplication(a, b int) int {
	return a * b
}

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("second divider can't be 0")
	}

	return a / b, nil
}

func Calculate(a int, b int, operation string) (int, error) {
	var result int
	var err error
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		return result, fmt.Errorf("operation can be one of these:+ - / *")
	}
	switch operation {
	case "+":
		result = Addition(a, b)
	case "-":
		result = Subtraction(a, b)
	case "*":
		result = Multiplication(a, b)
	case "/":
		result, err = Division(a, b)
	}
	if err != nil {
		return result, err
	}
	return result, nil
}
