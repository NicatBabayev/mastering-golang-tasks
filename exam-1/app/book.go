package app

import "fmt"

type Book struct {
	BookID int
	Title  string
	Author string
	Status string
}

func BorrowBook(userID int, bookID int, db *DB) error {
	userIndex, user := db.GetUserByID(userID)
	bookIndex, book := db.GetBookByID(bookID)
	if userIndex == -1 {
		return fmt.Errorf("requested user doesn't exist.")
	}
	if bookIndex == -1 {
		return fmt.Errorf("requested book doesn't exist.")
	}
	if user.BorrowLimit < 0 || user.BorrowLimit > 3 {
		return fmt.Errorf("user \"%s\" reached borrow limit.", user.Name)
	}
	if book.Status == "borrowed" {
		return fmt.Errorf("requested book \"%s\" already booked by another user.", book.Title)
	}
	db.users[userIndex].BorrowedBooks = append(db.users[userIndex].BorrowedBooks, bookID)
	db.users[userIndex].BorrowLimit++
	db.books[bookIndex].Status = "borrowed"
	return nil
}

func ReturnBook(userID int, bookID int, db *DB) error {
	userIndex, _ := db.GetUserByID(userID)
	bookIndex, book := db.GetBookByID(bookID)
	bIndexInBorrowedBooks := -1
	if userIndex == -1 {
		return fmt.Errorf("requested user doesn't exist.")
	}
	if bookIndex == -1 {
		return fmt.Errorf("requested book doesn't exist.")
	}
	if book.Status == "available" {
		return fmt.Errorf("requested book \"%s\" has not booked yet.", book.Title)
	}
	for i, j := range db.users[userIndex].BorrowedBooks {
		if j == bookID {
			bIndexInBorrowedBooks = i
		}
	}

	db.users[userIndex].BorrowedBooks = append(db.users[userIndex].BorrowedBooks[:bIndexInBorrowedBooks], db.users[userIndex].BorrowedBooks[bIndexInBorrowedBooks+1:]...)
	db.users[userIndex].BorrowLimit--
	db.books[bookIndex].Status = "available"

	return nil
}
