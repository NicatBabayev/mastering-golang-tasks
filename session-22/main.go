package main

import (
	"fmt"
	"session-22/app"
)

func main() {
	err := app.Init()
	if err != nil {
		fmt.Println("App Error:", err)
	}
}
