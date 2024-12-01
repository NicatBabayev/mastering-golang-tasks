package cmd_interface

import (
	"fmt"
	"os"
)

func commandHello(arg string) {
	fmt.Printf("Hello, %s!\n", arg)
}
func commandVersion() {
	fmt.Println("Version 1.0.0")
}
func commandHelp() {
	fmt.Println("Usage:")
	fmt.Println("hello <name>    - Greet the user.")
	fmt.Println("version         - Print the version of the program.")
	fmt.Println("help            - Show usage instructions.")
}

func handleSubCommands() {
	args := os.Args[1:]

	switch args[0] {
	case "hello":
		commandHello(args[1])
	case "version":
		commandVersion()
	case "help":
		commandHelp()
	}
}

func RunTask1() {
	handleSubCommands()
}
