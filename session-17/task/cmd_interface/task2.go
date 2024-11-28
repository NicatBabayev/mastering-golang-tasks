package cmd_interface

import (
	"fmt"
	"os"
	"strconv"
)

func sumInt(str1, str2 string) int {
	num1, _ := strconv.Atoi(str1)
	num2, _ := strconv.Atoi(str2)
	return num1 + num2

}

func handleTask2(arguments ...interface{}) error {
	var res interface{}
	args := os.Args[1:]
	if args[0] == "add" {
		res = sumInt(arguments[0].(string), arguments[1].(string))
	} else {
		res = arguments[0].(string) + arguments[1].(string)
	}
	fmt.Printf("Result: %v\n", res)
	return nil
}

func RunTask2() {
	args := os.Args[1:]

	if len(args[1:]) != 2 {
		fmt.Println(fmt.Errorf("Error: Invalid number of arguments for %s. Expected %d, got %d.\n", args[0], 2, len(args[1:])))
		os.Exit(0)
	}
	if args[0] != "concat" && args[0] != "add" {
		fmt.Println(fmt.Errorf("Only 'add' and 'concat' subcommands supported"))
		os.Exit(0)
	}
	err := handleTask2(os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println(err)
	}
}
