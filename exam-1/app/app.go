package app

import (
	"fmt"
	"os"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/inancgumus/screen"
)

func Run() {
	db := DB{}
	box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", TitlePos: "Top", AllowWrapping: false, Type: "Classic", Color: "Green"})
	screen.Clear()
	DisplayWelcomeMessage()
	LoadInitialData(&db)
	DisplayData(&db)
	time.Sleep(time.Second * 1)
	for {
		screen.MoveTopLeft()
		screen.Clear()
		DisplayWelcomeMessage()
		box.Print("Commands", ReturnCommandsMenu())
		commandOption := GetCommand()
		CommandRunner(commandOption, &db)
		time.Sleep(time.Second * 3)
	}
	// TODO Implement borrow procedure
	// TODO Implement return procedure
	// Implement available books listing
	// Implement borrowed books listing
}

func CommandRunner(commandOption int, db *DB) {
	switch commandOption {
	case 1:
		listAllBooks(db)
	case 2:
		listAvailableBooks(db)
	case 3:
		userID := selectUser(db)
		fmt.Println(userID)
		books, _ := returnBorrowedBooks(userID, db)
		DisplayBooks(books)
	case 4:
		listAllUsers(db)
	case 5:
		BorrowBook(db)
	case 6:
		ReturnBook(db)
	case 7:
		screen.Clear()
		DisplayWelcomeMessage()
		DisplayData(db)
	case 8:
		os.Exit(0)
	default:
	}
}

func GetCommand() int {
	var commandNumber int
	/* box := box.New(box.Config{Px: 2, Py: 0, ContentAlign: "Left", TitlePos: "Inside", AllowWrapping: false, Type: "Classic", Color: "Green"})
	box.Println("", "Please select operation by writing according number: ") */
	fmt.Print("Please select operation by writing according number: ")
	_, err := fmt.Scan(&commandNumber)
	if err != nil {
		fmt.Println(fmt.Errorf("input error: %w", err))
	}
	return commandNumber
}
func ReturnBook(db *DB) {}
func BorrowBook(db *DB) {
	userID := selectUser(db)
	/* 	availableBooks := returnAvailableBooks(db) */
	bookID := selectBook(db)
	borrowBook(userID, bookID, db)
	listAllBooks(db)
}
func DisplayBooks(books []Book) {
	fmt.Println("Books: ")
	for _, j := range books {
		fmt.Printf("BookID: %d;\t Title: %s;\t Author: %s;\t Status: %s;\n", j.BookID, j.Title, j.Author, j.Status)
	}
}
