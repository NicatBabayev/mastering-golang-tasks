package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"session-22/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type BooksHandler struct {
	db *gorm.DB
}

func NewBooksHandler(db *gorm.DB) *BooksHandler {
	return &BooksHandler{db: db}
}

func (h *BooksHandler) BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getBooksHandler(w, r)
	case "POST":
		h.postBooksHandler(w, r)
	case "PUT":
		h.putBooksHandler(w, r)
	case "DELETE":
		h.deleteBooksHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handlers
func (h *BooksHandler) getBooksHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathParts := strings.Split(path[1:], "/")
	// Filter out empty strings
	var filteredParts []string
	for _, part := range pathParts {
		if part != "" {
			filteredParts = append(filteredParts, part)
		}
	}
	if len(filteredParts) == 1 {
		books, _ := h.getAllBooks()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(books); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		id, _ := strconv.Atoi(pathParts[1])
		book := h.getBookByID(id)
		defBook := model.BookResponse{}
		if *book == defBook {
			http.Error(w, "Requested book not found", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(book); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
func (h *BooksHandler) postBooksHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathParts := strings.Split(path[1:], "/")
	var book *model.BookRequest
	var filteredParts []string
	for _, part := range pathParts {
		if part != "" {
			filteredParts = append(filteredParts, part)
		}
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &book); err != nil {
		http.Error(w, "Failed to decode JSON object", http.StatusBadRequest)
		return
	}
	if len(filteredParts) == 1 {
		err := h.addBook(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		id, _ := strconv.Atoi(pathParts[1])
		err := h.addBookByID(id, book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
func (h *BooksHandler) deleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathParts := strings.Split(path[1:], "/")

	id, _ := strconv.Atoi(pathParts[1])

	err := h.deleteBookByID(id)
	if err != nil {
		http.Error(w, "Failed to delete a book", http.StatusInternalServerError)
		return
	}
}
func (h *BooksHandler) putBooksHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathParts := strings.Split(path[1:], "/")
	var book *model.BookRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &book); err != nil {
		http.Error(w, "Failed to decode JSON object", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(pathParts[1])
	err = h.updateBookByID(id, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// Handle Methods
func (h *BooksHandler) findBookID(book *model.BookRequest) (int, error) {
	var bookResult model.BookResponse
	if err := h.db.Where("title = ?", book.Title).First(&bookResult).Error; err != nil {
		return -1, err
	}
	return int(bookResult.ID), nil
}
func (h *BooksHandler) getAllBooks() (*[]model.BookResponse, error) {
	var books []model.BookResponse
	err := h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("book").Find(&books).Error; err != nil {
			return fmt.Errorf("getBooksHandler error: %w\n", err)
		}
		return nil
	})
	if err != nil {
		return &books, fmt.Errorf("getBooksHandler transaction err: %w", err)
	}
	return &books, nil
}
func (h *BooksHandler) getBookByID(id int) *model.BookResponse {
	var resBook *model.BookResponse

	h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("book").Where("id = ?", id).First(&resBook).Error; err != nil {
			return err
		}
		return nil
	})

	return resBook
}
func (h *BooksHandler) addBook(book *model.BookRequest) error {
	err := h.db.Transaction(func(tx *gorm.DB) error {
		tx.Table("book").Create(&book)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (h *BooksHandler) addBookByID(id int, book *model.BookRequest) error {
	err := h.db.Transaction(func(tx *gorm.DB) error {
		book.ID = uint(id)
		tx.Table("book").Create(&book)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (h *BooksHandler) deleteBookByID(id int) error {
	err := h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("book").Delete(&model.BookRequest{}, id).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
func (h *BooksHandler) updateBookByID(id int, book *model.BookRequest) error {
	var currentBook *model.BookResponse
	currentBook = h.getBookByID(id)
	currentBook.Title = book.Title
	currentBook.Author = book.Author
	currentBook.PublishedYear = book.PublishedYear
	currentBook.Price = book.Price
	err := h.db.Transaction(func(tx *gorm.DB) error {
		tx.Save(currentBook)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
