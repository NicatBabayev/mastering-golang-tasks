package flags_arguments

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func RunTask4() {
	fs := flag.NewFlagSet("math", flag.ExitOnError)

	operation := fs.String("operation", "Default Operation", "Choose one of four simple math operations(+,-,*,/)")

	fs.Parse(os.Args[2:])
	num1, _ := strconv.Atoi(os.Args[4])
	num2, _ := strconv.Atoi(os.Args[5])

	if len(os.Args[2:]) != 4 {
		fmt.Println("Your input is invalid. Please write one operation and 2 numbers")
	}

	switch *operation {
	case "add":
		fmt.Println("Result:", num1+num2)
	case "substract":
		fmt.Println("Result:", num1-num2)
	case "multiply":
		fmt.Println("Result:", num1*num2)
	case "divide":
		if num2 == 0 {
			fmt.Println("Second number can't be 0 in division")
			os.Exit(0)
		}
		fmt.Printf("Result:%.2f\n", float64(num1)/float64(num2))
	}
}
