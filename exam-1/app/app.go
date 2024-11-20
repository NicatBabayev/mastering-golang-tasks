package app

import (
	"fmt"
	"os"
	"time"

	"github.com/inancgumus/screen"
)

func Run() {
	command := -1
	screen.Clear()
	db := DB{}
	LoadInitialData(&db)
	DisplayWelcomeScreen(&db)
	time.Sleep(time.Second * 2)
	for {
		DisplayCommandsMenu()
		fmt.Print("Please select the command: ")
		fmt.Scanln(&command)
		CommandRunner(command, &db)
	}

}
func CommandRunner(commandOption int, db *DB) {
	switch commandOption {
	case 1:
		DisplayAllBooks(db)
		time.Sleep(time.Second * 3)
	case 2:
		DisplayAvailableBooks(db)
		time.Sleep(time.Second * 3)
	case 3:
		DisplayAllBorrowedBooks(db)
		time.Sleep(time.Second * 3)
	case 4:
		DisplayAllUsers(db)
		time.Sleep(time.Second * 3)
	case 5:
		DisplayBorrowBook(db)
	case 6:
		DisplayReturnBook(db)
	case 7:
		DisplayWelcomeScreen(db)
		time.Sleep(time.Second * 3)
	case 8:
		os.Exit(0)
	default:
		fmt.Println("Please select correct option (1-8).")
		time.Sleep(time.Second * 3)
		return
	}
}
