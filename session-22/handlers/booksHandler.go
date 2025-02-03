package handlers

import (
	"net/http"
	"session-22/model"
	"strings"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBooksHandler(w, r)
	case "POST":
		postBooksHandler(w, r)
	case "PUT":
		putBooksHandler(w, r)
	case "DELETE":
		deleteBooksHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
}
func postBooksHandler(w http.ResponseWriter, r *http.Request)   {}
func putBooksHandler(w http.ResponseWriter, r *http.Request)    {}
func deleteBooksHandler(w http.ResponseWriter, r *http.Request) {}

func findBookID(book *model.BookRequest) (int, error) {}

func getBookByID(id int) *model.BookResponse {}
func getAllBooks() *[]model.BookResponse     {}

func addBook(book *model.BookRequest) error {}
func deleteBookByID(id int) error           {}

func updateBookByID(id int, book *model.BookRequest) error {}
