package intro_net_http

import (
	"fmt"
	"net/http"
)

func handleRoot(rw http.ResponseWriter, req *http.Request) {
	_, err := rw.Write([]byte("Hello World!\n"))
	_ = req // Make the static code analyzer happy
	if err != nil {
		fmt.Println(err)
	}
}

func RunTask1() {
	http.HandleFunc("/", handleRoot)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
