package user_inputs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunTask6() {
	var text []string
	fmt.Println("Enter text (type STOP to end): ")

	reader := bufio.NewReader(os.Stdin)

	for {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		if strings.ToLower(str) == "stop" {
			break
		}
		text = append(text, str)

	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("Captured lines in reverse order:")
	for i := len(text) - 1; i >= 0; i-- {
		fmt.Println(text[i])
	}
}
