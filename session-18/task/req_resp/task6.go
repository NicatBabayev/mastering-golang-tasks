package req_resp

import (
	"fmt"
	"net/http"
	"strconv"
)

func handleDivide(rw http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	a, _ := strconv.Atoi(params.Get("a"))
	b, _ := strconv.Atoi(params.Get("b"))
	var res []byte
	if b == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(rw, "400 Bad Request Division by zero is not allowed.")
		if err != nil {
			fmt.Println("Fprintf err", err)
		}
	} else {
		res = []byte(strconv.FormatFloat(float64(a)/float64(b), 'f', 2, 64))
		_, err := rw.Write(res)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func RunTask6() {
	task6 := http.NewServeMux()
	task6.HandleFunc("/divide", handleDivide)

	err := http.ListenAndServe(":8086", task6)
	if err != nil {
		fmt.Println(err)
	}
}
