package req_resp

import (
	"fmt"
	"net/http"
)

func handleGreet(rw http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	name := params.Get("name")
	var res string

	if len(name) == 0 {
		res = "Hello, Stranger!"
	} else {
		res = fmt.Sprintf("Hello, %s!", name)
	}
	_, err := rw.Write([]byte(res))
	if err != nil {
		fmt.Println(err)
	}
}

func RunTask5() {
	task5 := http.NewServeMux()
	task5.HandleFunc("/greet", handleGreet)

	err := http.ListenAndServe(":8085", task5)
	if err != nil {
		fmt.Println(err)
	}
}
