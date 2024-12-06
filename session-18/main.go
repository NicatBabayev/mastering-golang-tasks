package main

import (
	"fmt"
	"session-18/task/intro_net_http"
	"session-18/task/req_resp"
	"session-18/task/web_srv"
	"time"
)

func main() {

	//// Task 1
	fmt.Println("// Task 1.\nAddress: http://localhost:8081/")
	go intro_net_http.RunTask1()

	fmt.Println()

	// Task 2
	fmt.Println("// Task 2.\nAddress: http://localhost:8082/")
	go intro_net_http.RunTask2()

	fmt.Println()

	// Task 3
	fmt.Println("// Task 3\nAddress: http://localhost:8083/")
	go web_srv.RunTask3()

	fmt.Println()

	// Task 4
	fmt.Println("// Task 4\nAddress: http://localhost:8084/log")
	go web_srv.RunTask4()

	fmt.Println()

	// Task 5
	fmt.Println("// Task 5\nAddress: http://localhost:8085/greet")
	go req_resp.RunTask5()

	fmt.Println()

	// Task 6
	fmt.Println("// Task 6\nAddress: http://localhost:8086/divide")
	go req_resp.RunTask6()

	time.Sleep(time.Second * 360)

}
