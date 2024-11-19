package app

import "fmt"

func LoadInitialData(db *DB) {
	books := []Book{
		{
			BookID: 1,
			Title:  "The Go Programming Guide",
			Author: "John Doe",
			Status: "available",
		},
		{
			BookID: 2,
			Title:  "Concurrency in Practice",
			Author: "Jane Smith",
			Status: "available",
		},
		{
			BookID: 3,
			Title:  "Microservices Design",
			Author: "Alan Turing",
			Status: "available",
		},
		{
			BookID: 4,
			Title:  "Data Structures",
			Author: "Grace Hopper",
			Status: "available",
		},
		{
			BookID: 5,
			Title:  "The Art of Computer Prog",
			Author: "Donald Knuth",
			Status: "available",
		},
		{
			BookID: 6,
			Title:  "Clean Code",
			Author: "Robert Martin",
			Status: "available",
		},
		{
			BookID: 7,
			Title:  "Algorithms Unlocked",
			Author: "Thomas Gormen",
			Status: "available",
		},
		{
			BookID: 8,
			Title:  "Software Architescture",
			Author: "Martin Fowler",
			Status: "available",
		},
		{
			BookID: 9,
			Title:  "System Design",
			Author: "Scott Meyers",
			Status: "available",
		},
		{
			BookID: 10,
			Title:  "Effective Go",
			Author: "Rob Pike",
			Status: "available",
		},
	}
	users := []User{
		{
			UserID:        1,
			Name:          "Ali",
			BorrowedBooks: []int{},
			BorrowLimit:   3,
		},
		{
			UserID:        2,
			Name:          "Babek",
			BorrowedBooks: []int{},
			BorrowLimit:   3,
		},
		{
			UserID:        3,
			Name:          "Cabir",
			BorrowedBooks: []int{},
			BorrowLimit:   3,
		},
		{
			UserID:        4,
			Name:          "Diana",
			BorrowedBooks: []int{},
			BorrowLimit:   3,
		},
		{
			UserID:        5,
			Name:          "Elvira",
			BorrowedBooks: []int{},
			BorrowLimit:   3,
		},
	}
	for _, book := range books {
		err := db.addBook(book)
		if err != nil {
			fmt.Println(fmt.Errorf("load initial data error: %w", err))
		}
	}
	for _, user := range users {
		err := db.addUser(user)
		if err != nil {
			fmt.Println(fmt.Errorf("load initial data error: %w", err))
		}
	}

}
func DisplayData(db *DB) {
	fmt.Println("1. Available Books: ")
	for _, book := range db.books {
		fmt.Printf("  - \"%s\" by %s\n", book.Title, book.Author)
	}
	fmt.Println("2. Borrowed Books by Users: ")
	for _, user := range db.users {
		fmt.Printf("  - User %d (%s):\n", user.UserID, user.Name)
		if len(user.BorrowedBooks) != 0 {
			for _, borrowedBookID := range user.BorrowedBooks {
				book, _ := db.searchBookByID(borrowedBookID)
				fmt.Printf("    - \"%s\" by %s", book.Title, book.Author)
			}
		}
	}
	fmt.Println("End of Report.")
}
