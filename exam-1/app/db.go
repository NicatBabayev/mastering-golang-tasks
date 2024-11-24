package app

import "fmt"

type DB struct {
	books []Book
	users []User
}

// User related methods
func (db *DB) GetUserByID(userID int) (int, User) {
	index := -1
	var res User
	for i, user := range db.users {
		if user.UserID == userID {
			index = i
			res = user
		}
	}
	if index != -1 {
		return index, res
	} else {
		return -1, User{}
	}
}

func (db *DB) GetAllUsers() []User {
	return db.users
}

func (db *DB) AddUser(user User) error {
	index, _ := db.GetUserByID(user.UserID)
	if index != -1 {
		return fmt.Errorf("requested user already exists in database.")
	}
	db.users = append(db.users, user)
	return nil
}

func (db *DB) RemoveUser(user User) error {
	index, _ := db.GetUserByID(user.UserID)

	if index == -1 {
		return fmt.Errorf("requested user doesn't exist in database.")
	}
	db.users = append(db.users[:index], db.users[index+1:]...)
	return nil
}

// Book related methods
func (db *DB) GetAllBooks() []Book {
	return db.books
}

func (db *DB) GetAllBorrowedBooks() []Book {
	borrowedBooks := make([]Book, 0)
	for _, book := range db.books {
		if book.Status == "borrowed" {
			borrowedBooks = append(borrowedBooks, book)
		}
	}
	return borrowedBooks
}

func (db *DB) GetAvailableBooks() []Book {
	availableBooks := make([]Book, 0)
	for _, book := range db.books {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (db *DB) GetBorrowedBooksByUserID(userID int) []int {
	index, user := db.GetUserByID(userID)
	if index != -1 {
		return user.BorrowedBooks
	}
	return []int{}
}

func (db *DB) GetBookByID(bookID int) (int, Book) {
	index := -1
	var res Book
	for i, book := range db.books {
		if book.BookID == bookID {
			index = i
			res = book
		}
	}
	if index != -1 {
		return index, res
	} else {
		return -1, Book{}
	}
}
func (db *DB) AddBook(book Book) error {
	index, _ := db.GetBookByID(book.BookID)
	if index != -1 {
		return fmt.Errorf("requested book already exists in database.")
	}
	db.books = append(db.books, book)
	return nil
}
func (db *DB) RemoveBook(book Book) error {
	index, _ := db.GetBookByID(book.BookID)
	if index == -1 {
		return fmt.Errorf("requested book doesn't exist in database.")
	}
	db.books = append(db.books[:index], db.books[index+1:]...)
	return nil
}
