package web_srv

import (
	"fmt"
	"net/http"
)

func handleRoot(rw http.ResponseWriter, req *http.Request) {
	_ = req // Make the static code analyzer happy
	_, err := rw.Write([]byte("Welcome to hte homepage!"))
	if err != nil {
		fmt.Println(err)
	}
}
func handleAbout(rw http.ResponseWriter, req *http.Request) {
	_ = req // Make the static code analyzer happy
	_, err := rw.Write([]byte("This is the about page."))
	if err != nil {
		fmt.Println(err)
	}
}

func handleContact(rw http.ResponseWriter, req *http.Request) {
	_ = req // Make the static code analyzer happy
	_, err := rw.Write([]byte("Contact us at contact@example.com."))
	if err != nil {
		fmt.Println(err)
	}
}

func RunTask3() {
	task3 := http.NewServeMux()
	task3.HandleFunc("/", handleRoot)
	task3.HandleFunc("/about", handleAbout)
	task3.HandleFunc("/contact", handleContact)

	err := http.ListenAndServe(":8083", task3)
	if err != nil {
		fmt.Println(err)
	}
}
