package web_srv

import (
	"fmt"
	"net/http"
)

func handleLog(rw http.ResponseWriter, req *http.Request) {
	method := req.Method
	url := req.URL.Path
	fmt.Printf("Method: %s, URL: %s\n", method, url)
}

func RunTask4() {
	task4 := http.NewServeMux()
	task4.HandleFunc("/log", handleLog)

	err := http.ListenAndServe(":8084", task4)
	if err != nil {
		fmt.Println(err)
	}
}
