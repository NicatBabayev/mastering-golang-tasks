package app

import (
	"fmt"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/inancgumus/screen"
)

func FormatBorrowedBooksByUsers(db *DB) string {
	var res string
	for _, user := range db.users {
		if len(user.BorrowedBooks) > 0 {
			res += fmt.Sprintf("  - [%d] User %d (%s):\n", user.UserID, user.UserID, user.Name)
			for _, borrowedBookID := range user.BorrowedBooks {
				_, book := db.GetBookByID(borrowedBookID)
				res += fmt.Sprintf("      - \"%s\" by %s\n", book.Title, book.Author)
			}
		} else if len(user.BorrowedBooks) == 0 {
			res += fmt.Sprintf("  - [%d] User %d (%s): No Books borrowed\n", user.UserID, user.UserID, user.Name)

		}
	}
	return res
}

func FormatBorrowedBooksByUserID(userID int, db *DB) string {
	var res string
	_, user := db.GetUserByID(userID)
	for _, bookID := range user.BorrowedBooks {
		_, book := db.GetBookByID(bookID)
		res += fmt.Sprintf("  - [%d] \"%s\" by %s\n", book.BookID, book.Title, book.Author)
	}

	return res
}

func FormatAvailableBooks(db *DB) string {
	var res string
	books := db.GetAvailableBooks()
	for _, book := range books {
		res += fmt.Sprintf("  - [%d] \"%s\" by %s\n", book.BookID, book.Title, book.Author)
	}
	return res
}

func FormatAllBooks(db *DB) string {
	var res string
	books := db.GetAllBooks()
	for _, book := range books {
		res += fmt.Sprintf("- [%d] \"%s\" by %s (%s)\n", book.BookID, book.Title, book.Author, book.Status)
	}
	return res
}

func FormatAllUsers(db *DB) string {
	var res string
	users := db.users
	for _, user := range users {
		res += fmt.Sprintf("- [%d] User %d: %s\n", user.UserID, user.UserID, user.Name)
	}
	return res
}

func FormatBorrowedUsers(db *DB) string {
	var res string
	users := db.users
	for _, user := range users {
		if len(user.BorrowedBooks) > 0 {
			res += fmt.Sprintf("- [%d] User %d: %s\n", user.UserID, user.UserID, user.Name)

		}
	}
	return res
}

func DisplayAllBooks(db *DB) {
	screen.Clear()
	screen.MoveTopLeft()
	allBooks := FormatAllBooks(db)
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Inside", TitleColor: "Green", AllowWrapping: false, Type: "Classic", Color: "Green"})
	DisplayWelcomeMessage()
	box.Print("All Books", allBooks)
}

func DisplayAvailableBooks(db *DB) {
	screen.Clear()
	screen.MoveTopLeft()
	availableBooks := FormatAvailableBooks(db)
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Inside", TitleColor: "Green", AllowWrapping: false, Type: "Classic", Color: "Green"})
	DisplayWelcomeMessage()
	if len(db.GetAvailableBooks()) == 0 {
		box.Print("Available Books", "All books are borrowed")
	} else {
		box.Print("Available Books", availableBooks)
	}
}

func DisplayAllBorrowedBooks(db *DB) {

	screen.Clear()
	screen.MoveTopLeft()
	borrowedBooks := FormatBorrowedBooksByUsers(db)
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Inside", TitleColor: "Green", AllowWrapping: false, Type: "Classic", Color: "Green"})
	DisplayWelcomeMessage()
	if len(db.GetAllBorrowedBooks()) == 0 {
		box.Print("Borrowed Books", "No books borrowed yet")
	} else {
		box.Print("Users", borrowedBooks)
	}

}

func DisplayBorrowedBookByAllUsers(db *DB) {
	screen.Clear()
	screen.MoveTopLeft()
	u := FormatBorrowedBooksByUsers(db)
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Inside", TitleColor: "Green", AllowWrapping: false, Type: "Classic", Color: "Green"})
	DisplayWelcomeMessage()
	box.Print("Borrowed Books", u)
}

func DisplayBorrowedBooksByUserID(userID int, db *DB) {
	screen.Clear()
	screen.MoveTopLeft()
	u := FormatBorrowedBooksByUserID(userID, db)
	_, user := db.GetUserByID(userID)
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Inside", TitleColor: "Green", AllowWrapping: false, Type: "Classic", Color: "Green"})
	DisplayWelcomeMessage()
	box.Print(fmt.Sprintf("Borrowed books by user %s", user.Name), u)
}

func DisplayAllUsers(db *DB) {
	screen.Clear()
	screen.MoveTopLeft()
	users := FormatAllUsers(db)
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Inside", TitleColor: "Green", AllowWrapping: false, Type: "Classic", Color: "Green"})
	DisplayWelcomeMessage()
	box.Print("Users", users)
}

func DisplayBorrowedUsers(db *DB) {
	screen.Clear()
	screen.MoveTopLeft()
	users := FormatBorrowedUsers(db)
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Inside", TitleColor: "Green", AllowWrapping: false, Type: "Classic", Color: "Green"})
	DisplayWelcomeMessage()
	box.Print("Users who borrowed books", users)

}

func DisplayBorrowBook(db *DB) {
	userID := -1
	bookID := -1
	DisplayWelcomeMessage()
	DisplayAllUsers(db)
	fmt.Print("Please select the user which you want to borrow the book for: ")
	fmt.Scanln(&userID)
	DisplayAvailableBooks(db)
	fmt.Print("Please select book to borrow: ")
	fmt.Scanln(&bookID)
	BorrowBook(userID, bookID, db)
	DisplayBorrowedBookByAllUsers(db)
	time.Sleep(time.Second * 3)
	// TODO Add the condition and check if borrowed books are empty. If so write a message.

}

func DisplayReturnBook(db *DB) { //TODO Finish the method
	userID := -1
	bookID := -1
	DisplayWelcomeMessage()
	DisplayBorrowedUsers(db)
	fmt.Print("Please select the user which you want to return the book for: ")
	fmt.Scanln(&userID)
	DisplayBorrowedBooksByUserID(userID, db)
	fmt.Print("Please select book to return: ")
	fmt.Scanln(&bookID)
	/* 	BorrowBook(userID, bookID, db) */
	ReturnBook(userID, bookID, db)
	DisplayBorrowedBookByAllUsers(db)
	time.Sleep(time.Second * 3)
	// TODO Add the condition and check if borrowed books are empty. If so write a message.
}

func DisplayWelcomeMessage() {
	fmt.Println(" ███████╗    ██╗     ██╗██████╗ ██████╗  █████╗ ██████╗ ██╗   ██╗")
	fmt.Println(" ██╔════╝    ██║     ██║██╔══██╗██╔══██╗██╔══██╗██╔══██╗╚██╗ ██╔╝")
	fmt.Println(" █████╗█████╗██║     ██║██████╔╝██████╔╝███████║██████╔╝ ╚████╔╝ ")
	fmt.Println(" ██╔══╝╚════╝██║     ██║██╔══██╗██╔══██╗██╔══██║██╔══██╗  ╚██╔╝  ")
	fmt.Println(" ███████╗    ███████╗██║██████╔╝██║  ██║██║  ██║██║  ██║   ██║   ")
	fmt.Println(" ╚══════╝    ╚══════╝╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ")

}

func DisplayCommandsMenu() {
	screen.Clear()
	screen.MoveTopLeft()
	DisplayWelcomeMessage()
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Inside", TitleColor: "Green", AllowWrapping: false, Type: "Classic", Color: "Green"})
	commandsMenu := fmt.Sprintf("- [1] List All Books\n- [2] List Available Books\n- [3] List Borrowed Books\n- [4] List All Users\n- [5] Borrow Book\n- [6] Return Book\n- [7] Show Welcome Screen \n- [8] Quit\n")
	box.Print("Commands", commandsMenu)
}

func DisplayWelcomeScreen(db *DB) {
	screen.Clear()
	screen.MoveTopLeft()
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Inside", TitleColor: "Green", AllowWrapping: false, Type: "Classic", Color: "Green"})
	DisplayWelcomeMessage()
	box.Print("Welcome", "1. Available Books:\n"+FormatAvailableBooks(db)+"2. Borrowed Books by Users:\n"+FormatBorrowedBooksByUsers(db))
}
