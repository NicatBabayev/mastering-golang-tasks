package intro_net_http

import (
	"fmt"
	"net/http"
)

func handleRoot2(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("X-Custom-Header", "Learning Go")
	_, err := rw.Write([]byte("Hello World!\n"))
	_ = req // Make the static code analyzer happy
	if err != nil {
		fmt.Println(err)
	}
}

func RunTask2() {
	task2 := http.NewServeMux()
	task2.HandleFunc("/", handleRoot2)

	err := http.ListenAndServe(":8082", task2)
	if err != nil {
		fmt.Println(err)
	}
}
