package app

import (
	"fmt"
)

type Book struct {
	BookID int
	Title  string
	Author string
	Status string
}

func borrowBook(userID int, bookID int, db *DB) error {
	// Determine the user and the book
	user, userErr := db.searchUserByID(userID)
	book, bookErr := db.searchBookByID(bookID)
	if userErr != nil {
		return fmt.Errorf("user error: %w", userErr)
	}
	if bookErr != nil {
		return fmt.Errorf("book error: %w", bookErr)
	}
	// Check if user can reached the borrow limit
	if user.BorrowLimit > 3 && user.BorrowLimit <= 0 {
		return fmt.Errorf("user \"%s\" reached borrow limit.", user.Name)
	}
	if book.Status == "borrowed" {
		return fmt.Errorf("book \"%s\" is already borrowed.", book.Title)
	}
	// add book to user borrowedBooks
	db.users[userID].BorrowedBooks = append(db.users[userID].BorrowedBooks, bookID)
	db.users[userID].BorrowLimit++
	// change book status
	db.books[bookID].Status = "borrowed"

	return nil
}
func returnBook(userID int, bookID int, db *DB) error {
	user, userErr := db.searchUserByID(userID)
	book, bookErr := db.searchBookByID(bookID)
	if userErr != nil {
		return fmt.Errorf("returnBook user error: %w", userErr)
	}
	if bookErr != nil {
		return fmt.Errorf("returnBook book error: %w", bookErr)
	}
	var index int
	for i, j := range user.BorrowedBooks {
		if book.BookID == j {
			index = i
		}
	}
	// remove from the borrowedBooks
	db.users[userID].BorrowedBooks = append(db.users[userID].BorrowedBooks[:index], db.users[userID].BorrowedBooks[index+1:]...)
	db.users[userID].BorrowLimit--
	// change book status
	db.books[bookID].Status = "available"

	return nil
}
func listAllBooks(db *DB) {
	db.listAllBooks()
}
func listAvailableBooks(db *DB) {
	db.listAvailableBooks()
}
func returnAvailableBooks(db *DB) []Book {
	return db.returnAvailableBooks()
}
func returnBorrowedBooks(userID int, db *DB) ([]Book, error) {
	user, userErr := db.searchUserByID(userID)
	books := make([]Book, 0)
	if userErr != nil {
		return []Book{}, fmt.Errorf("viewBorrowedBooks user error: %w", userErr)
	}
	for _, j := range user.BorrowedBooks {
		book, _ := db.searchBookByID(j)
		books = append(books, book)
	}
	return books, nil
}
func selectBook(db *DB) int {
	var selectionID int
	fmt.Println("All Books: ")
	i := 0
	for _, book := range db.books {
		if book.Status == "available" {
			fmt.Printf("- [%d] %s.\n", i+1, book.Title)
			i++
		}
	}
	fmt.Print("Select the book: ")
	fmt.Scanln(&selectionID)
	res, _ := db.searchBook(db.books[selectionID-1])

	return res
}
