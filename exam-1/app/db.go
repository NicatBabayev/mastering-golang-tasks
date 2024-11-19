package app

import "fmt"

type DB struct {
	books []Book
	users []User
}

// Book related methods
func (db *DB) addBook(book Book) error {
	index, err := db.searchBook(book)
	if index != -1 {
		return fmt.Errorf("addBook error: %w", err)
	}
	db.books = append(db.books, book)
	return nil
}
func (db *DB) removeBook(book Book) error {
	index, err := db.searchBook(book)
	if index != -1 {
		db.books = append(db.books[:index], db.books[index+1:]...)
		return nil
	}
	return fmt.Errorf("removeUser err: %w", err)
}
func (db *DB) searchBook(book Book) (int, error) {
	for i, j := range db.books {
		if j.BookID == book.BookID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Requested book didn't found")
}
func (db *DB) searchBookByID(bookID int) (Book, error) {
	for _, j := range db.books {
		if j.BookID == bookID {
			return j, nil
		}
	}
	return Book{}, fmt.Errorf("Requested book didn't found.")
}
func (db *DB) listAllBooks() {
	fmt.Println("Books: ")
	for _, j := range db.books {
		fmt.Printf("BookID: %d;\t Title: %s;\t Author: %s;\t Status: %s;\n", j.BookID, j.Title, j.Author, j.Status)
	}
}
func (db *DB) listAvailableBooks() {
	fmt.Println("Books: ")
	for _, j := range db.books {
		if j.Status == "available" {
			fmt.Printf("BookID: %d;\t Title: %s;\t Author: %s;\t Status: %s;\n", j.BookID, j.Title, j.Author, j.Status)
		}
	}
}
func (db *DB) returnAvailableBooks() []Book {
	res := make([]Book, 0)
	for _, j := range db.books {
		if j.Status == "available" {
			res = append(res, j)
		}
	}
	return res
}
func (db *DB) listBorrowedBooks() {

	fmt.Println("Borrowed books: ")
	for _, j := range db.books {
		if j.Status == "borrowed" {
			fmt.Printf("BookID: %d;\t Title: %s;\t Author: %s;\t Status: %s;\t", j.BookID, j.Title, j.Author, j.Status)
		}
	}
}

// User related methods
func (db *DB) addUser(user User) error {
	index, err := db.searchUser(user)
	if index == -1 {
		db.users = append(db.users, user)
		return nil
	}
	return fmt.Errorf("addUser error: %w", err)

}
func (db *DB) removeUser(user User) error {
	index, err := db.searchUser(user)
	if index != -1 {
		db.users = append(db.users[:index], db.users[index+1:]...)
		return nil
	}
	return fmt.Errorf("removeUser err: %w", err)

}
func (db *DB) searchUser(user User) (int, error) {
	for i, j := range db.users {
		if j.UserID == user.UserID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Requested user didn't found")
}
func (db *DB) searchUserByID(userID int) (User, error) {
	for _, j := range db.users {
		if j.UserID == userID {
			return j, nil
		}
	}
	return User{}, fmt.Errorf("Requested user didn't found.")

}

func (db *DB) listAllUsers() {
	fmt.Print("Users: \n")
	for _, j := range db.users {
		fmt.Printf(" UserID: %d;\t Name: %s;\t BorrowedBooks: %v;\n ", j.UserID, j.Name, j.BorrowedBooks)
	}

}
