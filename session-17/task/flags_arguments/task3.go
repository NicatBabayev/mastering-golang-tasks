package flags_arguments

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func RunTask3() {
	fs := flag.NewFlagSet("greet", flag.ExitOnError)
	name := fs.String("name", "DefaultName", "Specify the name for to greet")
	uCase := fs.Bool("uppercase", false, "Use to convert the string to uppercase")
	err := fs.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
	}
	if *uCase {
		fmt.Println("HELLO,", strings.ToUpper(*name))
		os.Exit(0)
	}
	fmt.Println("Hello,", *name)
}
