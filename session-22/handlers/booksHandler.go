package handlers

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"session-22/model"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) BooksHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) getBooksHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Create and return JSON of the result

}
func (h *Handler) postBooksHandler(w http.ResponseWriter, r *http.Request)   {}
func (h *Handler) putBooksHandler(w http.ResponseWriter, r *http.Request)    {}
func (h *Handler) deleteBooksHandler(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) findBookID(book *model.BookRequest) (uint, error) {
	var bookResult model.BookResponse
	if err := h.db.Where("title = ?", book.Title).First(&bookResult).Error; err != nil {
		return -1, err
	}
	return bookResult.ID, nil
}

func (h *Handler) getBookByID(id int) *model.BookResponse {
	var resBook *model.BookResponse

	h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).First(&resBook).Error; err != nil {
			return err
		}
		return nil
	})
	return resBook
}
func (h *Handler) getAllBooks() (*[]model.BookResponse, error) {
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
	fmt.Println(books)
	return &books, nil
}

func (h *Handler) addBook(book *model.BookRequest) error {
	err := h.db.Transaction(func(tx *gorm.DB) error {
		tx.Table("book").Create(&book)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) deleteBookByID(id int) error {
	err := h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.BookRequest{}, id).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) updateBookByID(id int, book *model.BookRequest) error {
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
